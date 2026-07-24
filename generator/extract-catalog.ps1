<#
.SYNOPSIS
    Parse an exported ExchangeOnline.psm1 (the server-generated proxy module) with
    the PowerShell AST and emit a machine-readable JSON catalog of every cmdlet:
    name, parameter sets, and each parameter's type / attributes.

    This is the re-runnable front-end of the Go binding generator. When the API
    changes, re-fetch the psm1 (tools/capture/fetch-admin-spec.ps1) and re-run this.

.EXAMPLE
    pwsh ./generator/extract-catalog.ps1 -Psm1 .spec-cache/EXO-ExchangeOnline.psm1 -Out spec/catalog/EXO-catalog.json
#>
[CmdletBinding()]
param(
    [Parameter(Mandatory)][string]$Psm1,
    [Parameter(Mandatory)][string]$Out
)
$ErrorActionPreference = 'Stop'
New-Item -ItemType Directory -Force -Path (Split-Path $Out) | Out-Null

$tokens = $null; $errors = $null
$ast = [System.Management.Automation.Language.Parser]::ParseFile(
    (Resolve-Path $Psm1).Path, [ref]$tokens, [ref]$errors)

$funcs = $ast.FindAll({ param($n)
    $n -is [System.Management.Automation.Language.FunctionDefinitionAst] -and
    $n.Name -like 'script:*-*'
}, $true)

function Get-AttrArgText($a) {
    # positional args of an attribute, as source text (e.g. ValidateSet values)
    $a.PositionalArguments | ForEach-Object { $_.Extent.Text.Trim("'`"") }
}

$catalog = foreach ($f in $funcs) {
    $full = $f.Name -replace '^script:', ''
    $verb, $noun = $full -split '-', 2
    $defaultSet = $null
    # [CmdletBinding(DefaultParameterSetName='...')]
    foreach ($a in $f.Body.ParamBlock.Attributes) {
        if ($a.TypeName.GetReflectionType() -eq [System.Management.Automation.CmdletBindingAttribute] -or $a.TypeName.Name -eq 'CmdletBinding') {
            foreach ($na in $a.NamedArguments) {
                if ($na.ArgumentName -eq 'DefaultParameterSetName') { $defaultSet = $na.Argument.Extent.Text.Trim("'`"") }
            }
        }
    }

    $params = foreach ($p in $f.Body.ParamBlock.Parameters) {
        $pname = $p.Name.VariablePath.UserPath
        $type = 'System.Object'; $isSwitch = $false
        $sets = @(); $validate = @(); $aliases = @()
        foreach ($attr in $p.Attributes) {
            if ($attr -is [System.Management.Automation.Language.TypeConstraintAst]) {
                $type = $attr.TypeName.FullName
                if ($type -in 'switch','System.Management.Automation.SwitchParameter') { $isSwitch = $true; $type = 'switch' }
            }
            elseif ($attr -is [System.Management.Automation.Language.AttributeAst]) {
                switch ($attr.TypeName.Name) {
                    'Parameter' {
                        $set = [ordered]@{ name = '__AllParameterSets'; mandatory = $false; position = $null;
                            valueFromPipeline = $false; valueFromPipelineByPropertyName = $false }
                        foreach ($na in $attr.NamedArguments) {
                            $v = $na.Argument.Extent.Text.Trim("'`"")
                            switch ($na.ArgumentName) {
                                'ParameterSetName'                { $set.name = $v }
                                'Mandatory'                       { $set.mandatory = ($v -match 'true') -or (-not $na.Argument -is [System.Management.Automation.Language.ConstantExpressionAst]) }
                                'Position'                        { $set.position = [int]($v -replace '\D','') }
                                'ValueFromPipeline'               { $set.valueFromPipeline = $true }
                                'ValueFromPipelineByPropertyName' { $set.valueFromPipelineByPropertyName = $true }
                            }
                        }
                        $sets += $set
                    }
                    'ValidateSet' { $validate = Get-AttrArgText $attr }
                    'Alias'       { $aliases  = Get-AttrArgText $attr }
                }
            }
        }
        if (-not $sets) { $sets = @([ordered]@{ name = '__AllParameterSets'; mandatory = $false; position = $null; valueFromPipeline = $false; valueFromPipelineByPropertyName = $false }) }
        [ordered]@{ name = $pname; type = $type; isSwitch = $isSwitch;
                    parameterSets = $sets; validateSet = $validate; aliases = $aliases }
    }

    [ordered]@{ cmdlet = $full; verb = $verb; noun = $noun;
                defaultParameterSet = $defaultSet; parameters = @($params) }
}

$payload = [ordered]@{
    source        = Split-Path $Psm1 -Leaf
    cmdletCount   = @($catalog).Count
    generatedFrom = 'extract-catalog.ps1 (AST)'
    cmdlets       = @($catalog)
}
$payload | ConvertTo-Json -Depth 12 | Set-Content -Path $Out -Encoding utf8
Write-Host ("catalog: {0} cmdlets -> {1}" -f @($catalog).Count, $Out) -ForegroundColor Green

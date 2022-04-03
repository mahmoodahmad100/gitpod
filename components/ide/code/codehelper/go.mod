module github.com/gitpod-io/gitpod/code/codehelper

go 1.17

replace github.com/gitpod-io/gitpod/common-go => ../../../common-go // leeway

replace github.com/gitpod-io/gitpod/gitpod-protocol => ../../../gitpod-protocol/go // leeway

replace github.com/gitpod-io/gitpod/supervisor/api => ../../../supervisor-api/go // leeway

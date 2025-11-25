package gen

import (
	"github.com/zeromicro/go-zero/tools/goctl/model/sql/parser"
	"github.com/zeromicro/go-zero/tools/goctl/model/sql/template"
	"github.com/zeromicro/go-zero/tools/goctl/util"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
)

func genTag(table Table, field *parser.Field) (string, error) {
	if field.NameOriginal == "" {
		return field.NameOriginal, nil
	}

	text, err := pathx.LoadTemplate(category, tagTemplateFile, template.Tag)
	if err != nil {
		return "", err
	}

	output, err := util.With("tag").Parse(text).Execute(map[string]any{
		"field": field,
		"data":  table,
	})
	if err != nil {
		return "", err
	}

	return output.String(), nil
}

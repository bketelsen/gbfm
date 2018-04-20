package content

import (
	"fmt"
	"html/template"
	"strings"

	"github.com/gophersnacks/gbfm/models"
)

// selecterTplFuncused to build <select> form elements for a model that has a list
// of sub-models (i.e. and many_to_many relationship).
//
// sample usage to build a complete form element, including the
// surrounding <div class="form-element">:
//
//	<%= selecter(snack, topicsList, "Topics") %>
func selecterTplFunc(namer models.ModelNamer, list models.Lister, field string) template.HTML {
	optionsList := make([]string, list.Len())
	for i := 0; i < list.Len(); i++ {
		elt := list.EltAt(i)
		optionsList[i] = fmt.Sprintf(
			`<option id="%s">%s</option>`,
			elt.GetID().String(),
			elt.GetID().String(),
		)
	}
	options := strings.Join(optionsList, "\n")

	str := fmt.Sprintf(`
		<div class="form-group">
			<label class="active">%s</label>
			<select class="form-control" name="%s" id="%s-%s">
				%s
			</select>
		</div>
	`,
		field, // i.e. "Topics"
		field, // i.e. "Topics"
		strings.ToLower(namer.ModelName()), // i.e. "snack"
		field, // i.e. "Topics"
		options,
	)
	return template.HTML(str)
}

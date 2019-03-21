package controller

import(
    "github.com/firedial/midas-go/model"
)

func AttributeGet(attribute_name string, queries map[string][]string) []model.Attribute {
    ids, isExist := queries["id"]
    id := "";
    if isExist {
        id = ids[0]
    }
    attribute := model.GetAttribute(attribute_name, id)
    return attribute
}
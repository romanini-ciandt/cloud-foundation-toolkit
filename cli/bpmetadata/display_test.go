package bpmetadata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUIInputFromVariables(t *testing.T) {
	tests := []struct {
		name     string
		coreVars []BlueprintVariable
		UIinput  *BlueprintUIInput
	}{
		{
			name: "display metadata does not exist",
			coreVars: []BlueprintVariable{
				{
					Name: "test_var_1",
				},
				{
					Name: "test_var_2",
				},
				{
					Name: "test_var_3",
				},
			},
			UIinput: &BlueprintUIInput{},
		},
		{
			name: "display metadata exists and is in line with core metadata",
			coreVars: []BlueprintVariable{
				{
					Name: "test_var_1",
				},
				{
					Name: "test_var_2",
				},
				{
					Name: "test_var_3",
				},
			},
			UIinput: &BlueprintUIInput{
				DisplayVariables: map[string]*DisplayVariable{
					"test_var_1": {
						Name:    "test_var_1",
						Visible: true,
					},
					"test_var_2": {
						Name:    "test_var_2",
						Visible: true,
					},
					"test_var_3": {
						Name:    "test_var_3",
						Visible: true,
					},
				},
			},
		},
		{
			name: "display metadata exists and is not in line with core metadata",
			coreVars: []BlueprintVariable{
				{
					Name: "test_var_1",
				},
				{
					Name: "test_var_2",
				},
				{
					Name: "test_var_4",
				},
			},
			UIinput: &BlueprintUIInput{
				DisplayVariables: map[string]*DisplayVariable{
					"test_var_1": {
						Name:    "test_var_1",
						Visible: true,
					},
					"test_var_2": {
						Name:    "test_var_2",
						Visible: true,
					},
					"test_var_3": {
						Name:    "test_var_3",
						Visible: true,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buildUIInputFromVariables(tt.coreVars, tt.UIinput)
			for _, v := range tt.coreVars {
				dispVar := tt.UIinput.DisplayVariables[v.Name]
				assert.NotNil(t, dispVar)
				assert.Equal(t, v.Name, dispVar.Name)
			}

			assert.GreaterOrEqual(t, len(tt.UIinput.DisplayVariables), len(tt.coreVars))
		})
	}
}
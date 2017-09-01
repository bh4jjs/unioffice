// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package schemaLibrary_test

import (
	"testing"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/schemaLibrary"
)

func TestCT_SchemaConstructor(t *testing.T) {
	v := schemaLibrary.NewCT_Schema()
	if v == nil {
		t.Errorf("schemaLibrary.NewCT_Schema must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed schemaLibrary.CT_Schema should validate: %s", err)
	}
}
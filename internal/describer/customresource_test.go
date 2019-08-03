/*
Copyright (c) 2019 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: Apache-2.0
*/

package describer

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/vmware/octant/internal/testutil"
	"github.com/vmware/octant/pkg/store"
	storefake "github.com/vmware/octant/pkg/store/fake"
)

func Test_customResourceDefinitionNames(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	o := storefake.NewMockStore(controller)

	crd1 := testutil.CreateCRD("crd1.example.com")
	crd2 := testutil.CreateCRD("crd2.example.com")

	crdList := testutil.ToUnstructuredList(t, crd1, crd2)

	crdKey := store.Key{
		APIVersion: "apiextensions.k8s.io/v1beta1",
		Kind:       "CustomResourceDefinition",
	}
	o.EXPECT().HasAccess(gomock.Any(), gomock.Any(), "list").Return(nil)
	o.EXPECT().List(gomock.Any(), gomock.Eq(crdKey)).Return(crdList, nil)

	ctx := context.Background()
	got, err := CustomResourceDefinitionNames(ctx, o)
	require.NoError(t, err)

	expected := []string{"crd1.example.com", "crd2.example.com"}

	assert.Equal(t, expected, got)
}

func Test_customResourceDefinition(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	o := storefake.NewMockStore(controller)

	crd1 := testutil.CreateCRD("crd1.example.com")

	crdKey := store.Key{
		APIVersion: "apiextensions.k8s.io/v1beta1",
		Kind:       "CustomResourceDefinition",
		Name:       "crd1.example.com",
	}
	o.EXPECT().Get(gomock.Any(), gomock.Eq(crdKey)).Return(testutil.ToUnstructured(t, crd1), nil)

	name := "crd1.example.com"
	ctx := context.Background()
	got, err := CustomResourceDefinition(ctx, name, o)
	require.NoError(t, err)

	assert.Equal(t, crd1, got)
}

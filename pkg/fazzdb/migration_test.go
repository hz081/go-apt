package fazzdb

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMigration(t *testing.T) {
	t.Run("Migration", func(t *testing.T) {
		TestCreateTable(t)
	})
}

func TestCreateTable(t *testing.T) {
	require.Panics(t, func() {
		CreateTable("test_create", func(table *MigrationTable) {})
	})
}

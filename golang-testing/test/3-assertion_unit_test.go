package helper

/* Catatan

1. Assert dan Require sama seperti Fail() dan FailNow()
2. assert.Equal adalah method yang digunakan untuk mengetest
	apakah hasil function yang ditest sama seperti yang diharapkan atau tidak.
3. direkomendasikan menggunakan assert atau require ketimbang error biasa. karena
	menggunakan library testify pesan errornya sangat lengkap
*/

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAssert(t *testing.T) {
	result := Sum(2,2,3)

	// assert.Equal(testing.T, expected, actual, msg)
	assert.Equal(t, 6, result, "Expecting : 6")
	fmt.Println(result)
}

func TestRequire(t *testing.T) {
	result := Sum(2,2,3)

	// require.Equal(testing.T, expected, actual, msg)
	require.Equal(t, 6, result, "Expecting : 6")
	fmt.Println(result)
}
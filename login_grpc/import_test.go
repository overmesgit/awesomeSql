package login_grpc

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os/exec"
	"strings"
	"testing"
)

func TestImports(t *testing.T) {
	data, err := exec.Command("go", "list", "-json", "../login/").Output()
	if err != nil {
		log.Error(err, string(data))
		return
	}
	var M map[string]interface{}
	err = json.Unmarshal(data, &M)
	if err != nil {
		log.Error(err)
		return
	}
	imports := fmt.Sprint(M["Imports"])
	if strings.Contains(imports, "login_psql/models") {
		t.Fatalf("Denied import %v", imports)
	}
}

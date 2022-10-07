package keeper_test

import (
	"testing"

	"github.com/alice/checkers/x/checkers/types"
	"github.com/stretchr/testify/require"
)

// func get_addresses(name string) (string, bool) {
// 	var addresses_map map[string]string = make(map[string]string, 3)
// 	addresses_map["alice"] = "cosmos1jmjfq0tplp9tmx4v9uemw72y4d2wa5nr3xn9d3"
// 	addresses_map["bob"] = "cosmos1xyxs3skf3f4jfqeuv89yyaqvjc6lffavxqhc8g"
// 	addresses_map["carol"] = "cosmos1e0w5t53nrq7p66fye6c8p0ynyhf6y24l4yuxd7"

// 	if address, ok := addresses_map[name]; ok {
// 		return address, true
// 	}

// 	return "", false

// }

func getAddresses() (alice, bob, carol string) {
	alice = "cosmos1jmjfq0tplp9tmx4v9uemw72y4d2wa5nr3xn9d3"
	bob = "cosmos1xyxs3skf3f4jfqeuv89yyaqvjc6lffavxqhc8g"
	carol = "cosmos1e0w5t53nrq7p66fye6c8p0ynyhf6y24l4yuxd7"
	return
}

func TestCreateGame(t *testing.T) {
	msgServer, context := setupMsgServer(t)
	alice, bob, carol := getAddresses()

	createResponse, err := msgServer.CreateGame(context, &types.MsgCreateGame{
		Creator: alice,
		Black:   bob,
		Red:     carol,
	})

	require.Nil(t, err)
	require.EqualValues(t, types.MsgCreateGameResponse{}, *createResponse)
}

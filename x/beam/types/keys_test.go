package types

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBeamKey(t *testing.T) {
	beamID := GenerateSecureToken(8)
	key := GetBeamKey(beamID)
	require.NotEqual(t, beamID, string(key))
	parsedBeamID := BytesKeyToString(SplitBeamKey(key))
	require.Equal(t, beamID, parsedBeamID)
}

func TestOpenBeamsQueueKey(t *testing.T) {
	beamID := GenerateSecureToken(8)
	key := GetOpenBeamQueueKey(beamID)
	require.NotEqual(t, beamID, string(key))
	parsedBeamID := BytesKeyToString(SplitOpenBeamQueueKey(key))
	require.Equal(t, beamID, parsedBeamID)
}

func TestClosedBeamsQueueKey(t *testing.T) {
	beamID := GenerateSecureToken(8)
	key := GetClosedBeamQueueKey(beamID)
	require.NotEqual(t, beamID, string(key))
	parsedBeamID := BytesKeyToString(SplitClosedBeamQueueKey(key))
	require.Equal(t, beamID, parsedBeamID)
}

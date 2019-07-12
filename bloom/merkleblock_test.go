// Copyright (c) 2013-2016 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package bloom_test

import (
	"bytes"
	"encoding/hex"
	"testing"

	"github.com/btgsuite/btgd/chaincfg/chainhash"
	"github.com/btgsuite/btgd/wire"
	btcutil "github.com/btgsuite/btgutil"
	"github.com/btgsuite/btgutil/bloom"
)

func TestMerkleBlock3(t *testing.T) {
	blockStr := "0100000079cda856b143d9db2c1caff01d1aecc8630d30625d10e8b4b8b" +
		"0000000000000b50cc069d6a3e33e3ff84a5c41d9d3febe7c770fdcc96b" +
		"2c3ff60abe184f196300000000000000000000000000000000000000000" +
		"0000000000000000000000067291b4d4c86041b00000000000000000000" +
		"00000000000000000000000000000000000000000000000101000000010" +
		"00000000000000000000000000000000000000000000000000000000000" +
		"0000ffffffff08044c86041b020a02ffffffff0100f2052a01000000434" +
		"104ecd3229b0571c3be876feaac0442a9f13c5a572742927af1dc623353" +
		"ecf8c202225f64868137a18cdd85cbbb4c74fbccfd4f49639cf1bdc94a5" +
		"672bb15ad5d4cac00000000"
	blockBytes, err := hex.DecodeString(blockStr)
	if err != nil {
		t.Errorf("TestMerkleBlock3 DecodeString failed: %v", err)
		return
	}
	blk, err := btcutil.NewBlockFromBytes(blockBytes)
	if err != nil {
		t.Errorf("TestMerkleBlock3 NewBlockFromBytes failed: %v", err)
		return
	}

	f := bloom.NewFilter(10, 0, 0.000001, wire.BloomUpdateAll)

	inputStr := "63194f18be0af63f2c6bc9dc0f777cbefed3d9415c4af83f3ee3a3d669c00cb5"
	hash, err := chainhash.NewHashFromStr(inputStr)
	if err != nil {
		t.Errorf("TestMerkleBlock3 NewHashFromStr failed: %v", err)
		return
	}

	f.AddHash(hash)

	mBlock, _ := bloom.NewMerkleBlock(blk, f)

	wantStr := "0100000079cda856b143d9db2c1caff01d1aecc8630d30625d10e8b4b8b" +
		"0000000000000b50cc069d6a3e33e3ff84a5c41d9d3febe7c770fdcc96b" +
		"2c3ff60abe184f196300000000000000000000000000000000000000000" +
		"0000000000000000000000067291b4d4c86041b00000000000000000000" +
		"00000000000000000000000000000000000000000000000100000001b50" +
		"cc069d6a3e33e3ff84a5c41d9d3febe7c770fdcc96b2c3ff60abe184f19" +
		"630101"
	want, err := hex.DecodeString(wantStr)
	if err != nil {
		t.Errorf("TestMerkleBlock3 DecodeString failed: %v", err)
		return
	}

	got := bytes.NewBuffer(nil)
	err = mBlock.BtcEncode(got, wire.ProtocolVersion, wire.LatestEncoding)
	if err != nil {
		t.Errorf("TestMerkleBlock3 BtcEncode failed: %v", err)
		return
	}

	if !bytes.Equal(want, got.Bytes()) {
		t.Errorf("TestMerkleBlock3 failed merkle block comparison: "+
			"got %v want %v", got.Bytes(), want)
		return
	}
}

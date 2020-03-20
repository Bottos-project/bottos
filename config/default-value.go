﻿// Copyright 2017~2022 The Bottos Authors
// This file is part of the Bottos Chain library.
// Created by Rocket Core Team of Bottos.

//This program is free software: you can distribute it and/or modify
//it under the terms of the GNU General Public License as published by
//the Free Software Foundation, either version 3 of the License, or
//(at your option) any later version.

//This program is distributed in the hope that it will be useful,
//but WITHOUT ANY WARRANTY; without even the implied warranty of
//MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//GNU General Public License for more details.

//You should have received a copy of the GNU General Public License
// along with bottos.  If not, see <http://www.gnu.org/licenses/>.

/*
 * file description:  define constants for this blockchain
 * @Author: May Luo
 * @Date:   2017-12-01
 * @Last Modified by:
 * @Last Modified time:
 */

package config

// DEFAULT_BLOCK_INTERVAL define defalut interval of block production
const DEFAULT_BLOCK_INTERVAL uint32 = 3

//PRODUCER_TIME_OUT for actor to set time out with millseconds
const PRODUCER_TIME_OUT uint32 = 50

// BLOCKS_PER_ROUND define block num per round
const BLOCKS_PER_ROUND uint32 = 29

// VOTED_DELEGATES_PER_ROUND define voted delegates per round
const VOTED_DELEGATES_PER_ROUND uint32 = 28

// CONSENSUS_BLOCKS_PERCENT define consensus rate
const CONSENSUS_BLOCKS_PERCENT uint32 = 70

// MAX_DELEGATE_VOTES define max delegate votes
const MAX_DELEGATE_VOTES uint32 = 49

// DELEGATE_PATICIPATION define delegate paticipation
const DELEGATE_PATICIPATION uint32 = 33

// MAX_BLOCK_SIZE define max block size
const MAX_BLOCK_SIZE uint32 = 5242880 // Byte
// DEFALT_SLOT_CHECK_INTERVAL define default slot check interval
const DEFALT_SLOT_CHECK_INTERVAL = 500000

// BOTTOS_CONTRACT_NAME define system contract name
const BOTTOS_CONTRACT_NAME string = "bottos"

// BOTTOS_INIT_SUPPLY define bto total supply
const BOTTOS_INIT_SUPPLY uint64 = 1000000000

// BOTTOS_SUPPLY_MUL define dot num of bto
const BOTTOS_SUPPLY_MUL uint64 = 100000000

// TOTAL_DELEGATE_REWARD total delegate reward
const TOTAL_DELEGATE_REWARD uint64 = 64000000

// PRIVATE_KEY_REGEXP define private key format
const PRIVATE_KEY_REGEXP string = "[0-9a-zA-Z]{64}"

// PUBLIC_KEY_REGEXP define public key format
const PUBLIC_KEY_REGEXP string = "[0-9a-zA-Z]{130}"

// HASH_SHA256_LENGTH define sha256 hash value length
const HASH_SHA256_LENGTH int = 32

// DEFAULT_BLOCK_TIME_LIMIT define default block time limit when producing block
const DEFAULT_BLOCK_TIME_LIMIT uint64 = 1000

// DEFAULT_MAX_LIFE_TIME define max life time for a transaction
const DEFAULT_MAX_LIFE_TIME uint64 = 10000 //unit: second

// DEFAULT_MAX_PENDING_TRX_IN_POOL define max pending transaction num in local transaction pool
const DEFAULT_MAX_PENDING_TRX_IN_POOL uint64 = 30000

const DEFAULT_MAX_SYNC_DISTANCE_PUT_TRX_IN_CACHE uint64 = 3

// DEFAULT_OPTIONDB_NAME define default option db name
const DEFAULT_OPTIONDB_NAME string = "bottos"

// DEFAULT_OPTIONDB_TABLE_BLOCK_NAME define default option db table name of block
const DEFAULT_OPTIONDB_TABLE_BLOCK_NAME string = "Blocks"

// DEFAULT_OPTIONDB_TABLE_TRX_NAME define default option db table name of trx
const DEFAULT_OPTIONDB_TABLE_TRX_NAME string = "Transactions"

// DEFAULT_OPTIONDB_TABLE_ACCOUNT_NAME define default option db table name of account
const DEFAULT_OPTIONDB_TABLE_ACCOUNT_NAME string = "Accounts"

// DEFAUL_MAX_CONTRACT_DEPTH define max call contract depth
const DEFAUL_MAX_CONTRACT_DEPTH uint32 = 10

// DEFAUL_MAX_SUB_CONTRACT_NUM define max sub contract num
const DEFAUL_MAX_SUB_CONTRACT_NUM uint32 = 10

//DEFAULT_BLOCK_CONSENSUS_TIMEOUT define default block consensus timeout time 6 block 18s
const DEFAULT_BLOCK_CONSENSUS_TIMEOUT uint32 = 18
const DEFAUL_MAX_NOTICE_CONTRACT_NUM uint32 = 5

const PRIMARY_TRX_SESSION string = "primary"
const SUB_TRX_SESSION string = "subsession"
const CONSENSUS_VERSION string = "1.0"

const TRX_IN uint16 = 0
const TRX_IN_TTL uint16 = 1
const TRX_OUT uint16 = 1
const TRX_OUT_TTL uint16 = 16
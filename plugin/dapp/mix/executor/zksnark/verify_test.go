// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package zksnark

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeposit(t *testing.T) {
	verifyKey := "11d72c948d3d6b88a49d99b55ac035875fbecc3654a0e275bb731b2da3acb94610bdc3e5aea2738b230cac2988e1290cd711f9170808e1c8eee0b4d1682c663825c8f7d853ca655f502524091ce0a48e7229f05df8e40bd31839ba19ef9ea3b81d5d15428b59e5c1b523917a989c189a2b9ed799cb18318160b47829868138cb2b49730f693a497e022a14344374410004ed29ab3333a2c4ff5a63d486d841210969ea406d44fc0af8f5a3ded0fe76f9a90c6c42c518e3725d07f91c0bcda1330982f9732d7a542b11ef8c90c2bb39d9f85f8bc240c2bea852d4dda3961be1e22a33bf4a7120e37df97c58aa10ac88fc365f156bffad5f0709615a0389b6542401689987d98818e7c0c55ca092ae21dd9895a5baba579a68bac5807e4f37677e128265980bf81da9622334495acd490d986068d12a148ec26f6b16d74e22f19d2ac20e521154c443fe93dba26c2b1cb6c8fa2533b8029ba28c168b80398b2a980268c89af9b92ed851e5cb18f1a688ce8d654b7db547c6753516776e6a16662c14ef921c2f016514a09a2265112b8f24da3267828793c9e37a5fa156dd3f5e8b279ac92cb2f35e64e2620def2c45b51503ee4552a1da0cfb055711efc08d89fe0e3a407989f7b14a8b0353a5ca8907be1d2551b2c085b2104e3afba15e30ad0d0527b06bd267ab6dffdc80b5ba6ff86adfb8c3f87e48a4ab3ba46de238e400ca1c7f4a396c30d31b20c756a1789213657c0fff66d7e7f8b1a3ec9728a1aada9622b0c732c615c9c78322da5da67ab6861ee30eaba4b0103c956265310e743f65000000031b3e8a8ac3b0ff4f009474b7dfb2e99e8c1368440219b8418351d2de4ec193c61b8fad37c1d5028da1e6a6b67e2c6858868710598175f4c6c07850d5e8ab53c120380af328d225d79a6ad6a716b7421de1b6f26793683fb83fae14ab6858486d062af23453da53472221872353ae3ac2900f711ab285de1cd9562eed856333c425767d1fb4797b1a64836d343541193cbfcc0e748fbb8a8b2f315505927c20fe1a87bc97b5f5d4e61216b061f2daa0703510a9813d253f710d73ad862c916603"
	pubInput := "000000020e362c7578039fcdf533d7c4490cab0de14ad33f9aefa628582f0ce04c4660b8000000000000000000000000000000000000000000000000000000003b9aca00"
	proof := "2db78b7225b082beeafe17d3c1e485c95afddf64294fd9e5e5a8604c1caa27f42e0ec6996b900e06302db62480d55ba0ef78702c018833b14c15df7884dd4b721261a627b9aee7fb8e945bdcfe71f48e983951e84285e2c96d1f44488a0301821df565d7f619f6933e58aa6a3bcb674d1e61cf95dc78beb9fd352262fe663d0812d888d2e41525827eb98ff0f1134a8de72a12190fba0903a05ed8239b36e0f613bafd955cbc29e1c0f60a7051fc9c7f56377badefe0d8e0b84f707b2869885b042bba53f3a4c71621464e2b6d0726b6472f629a81727fe2eff57d05da6587920af479a38b2add7c2a5c93ce5265c823ce3fb262be7164751f6205b3931f73db"

	rst, err := Verify(verifyKey, proof, pubInput)
	assert.Nil(t, err)
	assert.Equal(t, true, rst)

}

func TestAuthorize(t *testing.T) {
	verifyKey := "1473ab4328616752c00ea248309e10bce4b92e3fd4e3047e30747512821e24741a59bb2115d794d53e8a123b284f59acfccfe5fd76c87a4ee5cf72a24e415d52049eddd126746ca46f1a6b5151354844fba1cc136f56dc8d52ee087e77ee545a24793c266c6760371f66e3ddcbd2f50fdb40ebcd63327d0daeccace522636aa42525af4490cdbcba010d5d60b17b899a5d4ce32a4f028153d18ccb9b54108e35215fc55fa11835b75755f4b2e71af4787c7d8424c8fcde9b59d0283fa1a6596c287a8695a5c4fbe45979ba25930455fddeec304c1e0fe1d980730a262bd3c15a097ed012586cc746756ff350ba08c0f7e97b1e3368ad99d4b413419544a5ceea09b9dbb381264e8a786e9875c7fc9989ad119aa3a8e1b07f170dfc29b6603ac1245629f9260508d08747a7b8ddc2502be73cdc16e467d47df69aa71575d39119197ce70dbe1e1f1d8f1b45a3dddc63e6e236794fe08b42080e6746e8e22628f70d436597422407240ed94d46b32d0914451225dfeb64124f7215476148092e190010e75b4d11104f80c5259fda8778d6886fcbf789c48a13e21d7357eafd8a8308b0ae3dfcc341e715bdab7de49fb699330025da6060eb91b5a79cbe25e1cb3705f350c2d348ef3df4a6cccb11e68874fc8e75d617b3934a7c654c8c854f3752091291b80aa18de8fbd41b0ec39ebbd96ca685e6386da6cc204ae13a9184d18b044f26507b4b443e60e02162492364452a83246e9643adbd42e87316dc04382b29c76acfba7614956b9d9aed7e49be9258268e22adbba0aee818a3f8bc1c3bf40000000411f40834f563e79b2ffabc45c24797299ee34346888f34802b784b62427bd5d429f5117115d7162353b65557850a4a2f99e9834ca8ad95830ede631d897048521fabfd0dd4246c4d8e98062daf771ded13921619e76df0ae190ed0200757684e02ef7e138bf53134af024817312c725cc8a6617c8ecdd66d65e90576f59596c90358b4ce57d7b4832639096c4fc814fa067b3d4b11a4ba8626c543b26b68bac505c2ac761fb39a0b160010b4085dbe7ab2f4c940b96fa75fbcd0fefa6fbf12811e86674dd28531b9b264e6f098b91344df5755dd28a41696d1ffa878f5dd9ada0be976bdfcb9f5538df6dc9eee17babeeecaf3c93b4f317d189bb82cb0b7c264"
	proof := "1d638e78621f04eb8e579b369825182c07f21c48b9d199326b25cd5a9b88f612026a31d84d53689bb794e88735eb2fbaa03bd1eccf2f58a718b2710ec3531b642e32fa1f41bf952784763a6d6f924f5f031af6dbfadd3435ba11ba83b0c245890e0df56b0cc2fdb40101327ebab62e7adbce8c85ca7cb61de9682156253f256926ed3107f9b15a36c053c91ac6bd8ca76c1601721d0f35a10f9127bd6a16e2ed25f429360d602c25b26bb2fdbaf2115dd1a8ba86bdd32200a22efbf492fb467128dab9f0d7b22c554b0453f2f237da4e050d7179410aa39b8cab71bdd60aba802f037709f65500b8669a83e4f1d6f1ddec35d6e0794d4dabfadd4da7bcd7c1c8"
	pubInput := "00000003135823693902f72fb4d0793e576bb43919c5b41e128dd17b7f5c72a48257c86d0a5732da4df6bdcfeb79b586c3452090fd1fb29acb8248b2c1788594fb38276b294492f64fb29b7b065de3cf177ae35fc8334dc5fbb5cf29cbb1c901b482f1a6"

	rst, err := Verify(verifyKey, proof, pubInput)
	assert.Nil(t, err)
	assert.Equal(t, true, rst)

}

func TestTransferInput(t *testing.T) {
	verifyKey := "153cd9b6c5bd77f41c63012def12cfa719cedbbb7c376353ada8865bd9d386a22ebdbbd1da14114ff44d1ffd49d3a1f108a6affd4614c9c519fc84425bec6293145cb2f2e7c8945ed7e9a7a7ef02c1555bb0886e7c15a8be254cbab1adbbe33919b40f19c82819808f1f87ac6ae96657a60ff37557ec55e73043f51a3ae17ade00973ce23495e3dc7941b468ff0e9b9b0ae73d722fc5401983a8def9f48ab76020b317cf6ccbc216c4a20210c34e061b2d3f74e8a1645b8be4b7ff073abdb22923d6b82003695d70bf025442860d9979f4b547569d1c79bcb36bd85498d66969304a7edfa12d4ef559b07b93834024f424cc73d115e231d31c0489d22a19403321583c42261bb161b3566c574c42bfc80ef00bb5890e40b90ef3361eadb38f181144271f45c753afda2f81bdd2868e1adb7b2519d2f75ddb05e7fe7250a2bfc00beb6540deba4e021b1edc5d0c5e08015674d5db33c591714297acf1f02c35af1069621ead725fc0527099915ea5de69282cb9850de96495e86b61670f9146b411a9ea9a4df35af15a333495475ddd34df53692471d09e7fbc337471829f64d203a199091300223531a0298b17ad923c29f8982b32862b1c091d970f9b5cc9ef23a7761a669a5dd9d631424ea8879b89b22afcee997fc6ab95cdf9dc9baac6461c1bc5f5736b339568767e9ccab43e63dbfbd3f3a94f1100f566fa5f72029b701fda13678ff8db802156bce08bcc2ededa2914ac45d4be6d459abd612fbdd77718874dcb119724bcd71c8de5e10f7c79f4b66434f8b9a1e2a174d92cb10b1ca70000000628e6eaab99088247e8106f52f47b5940f3c5294d2283006bf3c29532c59a4c5a0e931758722334e0299af3b6cb36e9457fcac0bcd5440d80355babd3c05e88211e1e7f5bcdc4863afa2d49c743652305f971c92920c0bb61e76a1bb51f08156303ac2e2f58c0a6f2627b820ff08f0cedbbbfd99fa840b79de57e0f6ebf5969ce2efddfb68b6c57535a6be061335009ea879087b15def360546c5f1a4b0f737bf0d785f07ada66b23c6a037ce372994fd072a97c16d74c513f63e987217ec8dc210cbcb5c6dd8ffe745a63e88470fc09cade0ef877c4f107c08070fa19674618813e5be4d5db40af84bd83db80fccadee83af1549c92f4c935256bd5dd68b97f32e8dff5346fb5fa554c9ce617774c2a43cd7b6c5d5edbc9cd0dce0ba74b5ec87043b59f5b19c20057780c99db499d2565bc9f90c252f27cc37b2f5c9c69f504c13357153d572be7823f67ca794ca52088e20396eb7bef837109f542041fdd56310c5bd398cbfa8e588fd8342cbd103f58249481f5af3eac33f15add8fe4a64e4"
	proof := "07a8a57c1a0880229bf23afcb53fcd28a86cfb6dc80c4c9d24c343217f37d4cf0033eab8a97b46c865a5113ddc304ef0f4f2b8d47dd406aefac7c5fef4b5d29229cf43b27a52a222b3b6af43674e931d13e28582938ad63258a4ffe38e6512b701cc88b20f919bc0532f1164fbf71a0d952ad73408772d82b39b9eace52d00450d00e83776239d2086a407e2e9386b51388a3bee6de4b190ef2cdaaa2a9ab0771b1380bf6b82f8156abeefbaf6697ba2338659a88f2e2db6fe20db597edcf6762fba5aaaead8a565666305df3b74f20b7c4932d3e4deba854120727938ad6eb0108800c2f787edb56d9b2cf3f95a49af2186ea66e1a2e59ba00492d4c881d370"
	pubInput := "00000005135823693902f72fb4d0793e576bb43919c5b41e128dd17b7f5c72a48257c86d294492f64fb29b7b065de3cf177ae35fc8334dc5fbb5cf29cbb1c901b482f1a62c8cb0adf8339476c3a526ecd7f4eb8ce8efd9ce9680259f4b8b0f789d1ae6d92140ea0dda62736f1f4648e658a7987534fab27de21e2622dc2ffa249e958dcc2904614f6173fb2b98656a63e2d41abda2daf6ed340a4fc03426962361b14107"

	rst, err := Verify(verifyKey, proof, pubInput)
	assert.Nil(t, err)
	assert.Equal(t, true, rst)

}

func TestTransferOutput(t *testing.T) {
	verifyKey := "13934cfcc6ff87d92bd9bf14fc22abdda42ecfd544a0dd8b5b38d8936cb91f7d2ab93a5f2e54a2b09ae0f822be44bb304d2e8a3a9a2590e134cef4de7ca61e6222f85e2b09a87301942fa12d8755467ce88b30679f703b48a3ed973b526648c614058f734069dc99be9312b0f64b46dbfe8b5a1c7a7b2abd301417abe61cd7580f07fb23e6b6e998dbc16fa3f5e9d0489ec2867f496f79364628731c8a81159d26ba7de29f451d5b6128224f398ad7797a605b022f5a6de1b59276a339d88921218971e393ea6027f758ddffe692b11269b5ab17f72266661cc452b83f83a70818d7a37dabff75b1995952a63c45ef51695e18ce1b9b6bf0da5dbe3c9625ad090dd507dc37e610ea95e32f2109bd7ae0783484675a445cf396ca0653d4e3c71b27c0b8d2d3239ce11da02406913fbec4d973bb1ff1922a6d5d295581f635afc22cc1379afb9f5158d88baed9020b8d393f1c7157b30a36724d6efcf162fad84e20cd61fa92f99ef48a4ab4a0d93f91f06204addb1686344476605f2089defb561c4b0f8502b3dbe27395b86e8535fc4ca14c25c071880f8153e931194473b2d809dae5160125311328f69840cc36c0d68e3b135827f79076ec6cde7c94b607792485812b35437227068c0122eaa92ab9e181a7b6776b681d3c6d58d3a2d7e8462537bc06071d29fc3ec2e789d343359fd2adb5aa1e343033b6456cd9f9c9859120c39266d6032c2e4cd703c2ebaf1ed66edeaa95b8e9363190cb2e33b7c430ad2e7d9a9587ed281bf44e2305aa20c3e031ede1cfc40cfc600b09b45b0647cf1100000004269eb69c9017e0e3e8f2fefae3d2bf072ce5a9369b9dc4a949863a096b8c2fe70e53a7aaf51d02342549315f63b45527dbc1fc8b640b471337ea89ab78179c6305ba8d54c55ef56feb4623a09629f70164572993ac8ff3472753d643f325622220f4e4183a52903095d6989f5f3aaa18da0693d815e18eacab419e1306a492d42b1a34104b599e2c91e697969f4e47fb2305f5e0e587b1616c2cf53d4ac5ebba004710a9c23bfd6dfe3b0b4ce9b934fc6a6b2304640d56fc2ecdc20d37dbe352048691666bb7323b37f75cb2a3c078e05ee3d38ba666f33ee3c9cf9cbe2e1a901fcfbf411bdbf8b3e671579c29fe2da9b29cc1e99ab30801d8b9a1bd87c5cd8c"
	proof := "2df673e3d7c309b45f592fbdb2cfbaf00c4d65c5d58d683f5b24750cab9c058100927629a9a0e2aa0f79942ae44bde300f4f840e7c5a84dc98944e5ec75bd978034b179b21e148bd797cf6b4aa731e4f573e9ae310ad31b043b3df4dfe1744c004cd996f719753d497bc1ca01d2c20c3a8dd9bfb62e1a489e40c7d3568f492d613caada98166f244580f152594df359a4392fd594350d473b1bd51ab71ae21871ebec5f68adcada007fa4e2ab3ab4c61cfd5a4bc6772f4095e0c87e1c5713e6f1c68e3687cce8f3fc51e4e788422ece75175234ef5a7ce7aae115447723a4e4c2f4c0bedd98836af4c0930dfdbfd716e32904641a3fc8465191ab48425881d23"
	pubInput := "000000030844619e8ce593a653511d0c6e34b0ef7afab9ca112f3a44de6a3723a1036d6128a8c530d58f624d4e1c02a601ab29ebee175d1e44c4686ca9cfac192869cdce1e11821c7b033c0fd49f14e04e6b42a546ff397b84103b755327afe6bb7939b1"

	rst, err := Verify(verifyKey, proof, pubInput)
	assert.Nil(t, err)
	assert.Equal(t, true, rst)

}

func TestTransferChange(t *testing.T) {
	verifyKey := "13934cfcc6ff87d92bd9bf14fc22abdda42ecfd544a0dd8b5b38d8936cb91f7d2ab93a5f2e54a2b09ae0f822be44bb304d2e8a3a9a2590e134cef4de7ca61e6222f85e2b09a87301942fa12d8755467ce88b30679f703b48a3ed973b526648c614058f734069dc99be9312b0f64b46dbfe8b5a1c7a7b2abd301417abe61cd7580f07fb23e6b6e998dbc16fa3f5e9d0489ec2867f496f79364628731c8a81159d26ba7de29f451d5b6128224f398ad7797a605b022f5a6de1b59276a339d88921218971e393ea6027f758ddffe692b11269b5ab17f72266661cc452b83f83a70818d7a37dabff75b1995952a63c45ef51695e18ce1b9b6bf0da5dbe3c9625ad090dd507dc37e610ea95e32f2109bd7ae0783484675a445cf396ca0653d4e3c71b27c0b8d2d3239ce11da02406913fbec4d973bb1ff1922a6d5d295581f635afc22cc1379afb9f5158d88baed9020b8d393f1c7157b30a36724d6efcf162fad84e20cd61fa92f99ef48a4ab4a0d93f91f06204addb1686344476605f2089defb561c4b0f8502b3dbe27395b86e8535fc4ca14c25c071880f8153e931194473b2d809dae5160125311328f69840cc36c0d68e3b135827f79076ec6cde7c94b607792485812b35437227068c0122eaa92ab9e181a7b6776b681d3c6d58d3a2d7e8462537bc06071d29fc3ec2e789d343359fd2adb5aa1e343033b6456cd9f9c9859120c39266d6032c2e4cd703c2ebaf1ed66edeaa95b8e9363190cb2e33b7c430ad2e7d9a9587ed281bf44e2305aa20c3e031ede1cfc40cfc600b09b45b0647cf1100000004269eb69c9017e0e3e8f2fefae3d2bf072ce5a9369b9dc4a949863a096b8c2fe70e53a7aaf51d02342549315f63b45527dbc1fc8b640b471337ea89ab78179c6305ba8d54c55ef56feb4623a09629f70164572993ac8ff3472753d643f325622220f4e4183a52903095d6989f5f3aaa18da0693d815e18eacab419e1306a492d42b1a34104b599e2c91e697969f4e47fb2305f5e0e587b1616c2cf53d4ac5ebba004710a9c23bfd6dfe3b0b4ce9b934fc6a6b2304640d56fc2ecdc20d37dbe352048691666bb7323b37f75cb2a3c078e05ee3d38ba666f33ee3c9cf9cbe2e1a901fcfbf411bdbf8b3e671579c29fe2da9b29cc1e99ab30801d8b9a1bd87c5cd8c"
	proof := "0ae1034ce02cfb3479e838e499ec765bfa79636fee38c78f2ff3372f989ae1e52c3aa7a9768b33b41298fe189b3c37e27580089f54e5fd7c26b67cf681ec8a630e6055badd6dda947a0acd763a1ba8486f970ead2f08ef929ccf59b41dbfcdf5176ef886fe5c10245d8e951c93cc2619def52b5d8824392684c4f7a27b1955e223ad64ae429ccda0870686434f58b919e6166d169a0296687d88fa73da88d0b20074e395a19761058ccf1e280fec7563e0940b203e80a401f276b94b94c8925105c9571a53618379329de48871d56fdc4775c8fc3b8924ff35c6c6221c69cfb923ebc77245d3bd7343e1912cccb66f535ef475a8ac352086bdb132d41d1c1bb1"
	pubInput := "00000003000c83c30ec50aa2a60e1755e24de97c5faeeda1556fe3d65cc7ef9cdc9a06db1ed770eed86c3f980b65cf5e1676da2ec33ba26b00b7fa23cba7a85428d54f491fc961717b7a71489627073e68573b8b97e6182f77a41b52ec630fca04d338ab"

	rst, err := Verify(verifyKey, proof, pubInput)
	assert.Nil(t, err)
	assert.Equal(t, true, rst)

}

func TestWithdraw(t *testing.T) {
	verifyKey := "17497f382733020c2441ee2179d468e4a871353ecf52e7bc7f57600703e6cd0417c8a91b86c1e6e55316bd1ba88db8c817b646add9c5289497faf705e2e0a50510a31a9fbbcbec3399e15f1caea1e372a51827e08335606f7f1a9428539e4a3025ea43142772f0300fcb6e2e03057be5deebbc7e00cad1019da0452c2ab90d582caa5e4c22103eadf15f5ded717187934e4e9a44997766d6a770c42ff520e00d1d5a65abd94ac973cfb60d07499505d4db9584ccb428041336610d422a6d203007171dc5d586c536b3f49c2daf4d3b0a095a67ea759b49f182c87b41ac03848722b838c495beb5537c43fa7175bc5346743a2eb620856d0daa92614f283027ae2aff76e2a53c74648f52d9bd65ccf3064f2b9d9230e45bdaaa2214aa593d90180e85e2e045eac57a2b2b86d3eb23215bdc6b8b0ecb80e37374d41e1f7101acb502a75d203e2e058258f53f8c58192c8f0849a67a5ab5c7d26d782e4e1a10669c03285a2eae35d76a09c20841c49e72e7d7896da7a788d146836d5ba60bd2981b0eeabf8ee62fd2406056b556db292d18cbda9bcbe008791cb96fcb17ca6a33b0090652012ee82cce78f6ba28fe807cd3f668ba1fbdf795002d5ca1335a5de65307936b649d057637f8ff815d86aef66fc49960d02a8244cbd64658bc5758ee2a271e7a651fd9ef47bd7683c5457c724365c4a49e5e4ab88c980e33292ca1467c1e4b36114f7ed29b48bc2cf422be5c8ca374647971eb7daa1016348546190a762cc190a0d413e00c81f098a84e880bb53024be29917ebba4a55d59c40a606c0d0000000519dd32ea5a02431cbd4c38143012542385d9085b4c5d0b4f176285feae96d232061ed041ec53ee1897dc1eb82f1456bcb939b7512fe60548c97873bc602606c91f1d2dea1819016ef4bf1f1f3a6e074f04f04117d0628375f15779bc612078a302aa853d4d2e348cf05422efffd585fea5826ae63770b35be4ba8087092cf5150560d8d5e21d5b60413658c69f52d0ab2b399490f7e572ad8612a08472721ceb107ffe99a3a4de489fa4849a63eb5a574ddac2ccf7afe89c06c77c43a94e4d4b1c1250247c5352bd7b34dcd075f52cb31c5f4c96ba3120355a3dc010d330c4e02707c707d9cc3772ce61f7ff8bbd50d2633453a21d10edf821ceea93ba205bee2955c296632546fffbbdfdeea2f1aa67a9b1a90903ea12a3818d5e44f9964a6109afe1cbce7b5bccdf71718f22a3d57f6afd1a7976a75c093410520cf79bebf8"
	proof := "00f1469944fe6c7ac8e22ad7b6154db6421aaf4f0bbfe9407d781e7a2d40f82728d9b5d89fc75554eef3ad7f2aa4a7a04cde429be37d6502e8ea828572dfe66d0c73de4eb15f95224e78e49761d8a8cdb10fd3749da6f31e456b368af832ea621028a13e268939d66da0debb182518b86fa228be125e482f1d8f4ced94d854b6140c8b098a1eabdc7b4b6f4bc656259c44727db08652c3a3d993941c64a4aa9b0955c817a597fa51369113da6d356e7888ce3dcb9effb8ca00f7d12775f01c371fbb255b9a9a6411a22d18a63b0f92133f757253224e0e16f00b4f4dc3e3029c130f2e29a81c8bcd8fff051b139652e79bb30b65153ddaa1b01e08ce3bcae02f"
	pubInput := "000000041bff1e3aa6fae545c2eeca6ee1c06aa5ac0560a20b79bc9fa67387b27812ec0500000000000000000000000000000000000000000000000000000000000000002efad90a52dbaf0ac425ca9cdb871d5c0499994e99ade0485cffbe579775a7010000000000000000000000000000000000000000000000000000000023c34600"

	rst, err := Verify(verifyKey, proof, pubInput)
	assert.Nil(t, err)
	assert.Equal(t, true, rst)

}
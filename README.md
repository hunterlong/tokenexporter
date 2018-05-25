# TOKENexporter  [![Docker Build Status](https://img.shields.io/docker/build/hunterlong/tokenexporter.svg)](https://hub.docker.com/r/hunterlong/tokenexporter/)

A lightweight Prometheus exporter that will output ERC20 Token balances from a list of addresses you specify. TOKENexporter attaches to a geth server to fetch token wallet balances for your Grafana dashboards. You can also use [BTCexporter](https://github.com/hunterlong/btcexporter) for Bitcoin balances and [ETHexporter](https://github.com/hunterlong/ethexporter) for Ethereum balances.

## Watch Addresses
The `addresses.txt` file holds all the addresses to fetch balances for. Use the format `name:address` on each new line.
```
etherdelta:0x8d12A197cB00D4747a1fe03395095ce2A5CC6819
bittrex:0xFBb1b73C4f0BDa4f67dcA266ce6Ef42f520fBB98
```

## Running the Exporter
You can easily run this ERC20 Token balance prometheus exporter with the docker command:
```
docker run -it -d -p 9015:9015 \
  -v /myfolder/addresses.txt:/app/addresses.txt \ 
  hunterlong/tokenexporter
```

## List of all ERC20 Token Addresses
TOKENexporter uses [ethTokens.json](https://github.com/kvhnuke/etherwallet/blob/mercury/app/scripts/tokens/ethTokens.json) to get all the ERC20 contract addresses that your account may be holding.

## Build Docker Image
Clone this repo and then follow the simple steps below!

##### Build Docker Image
`docker build -t hunterlong/tokenexporter:latest .`

##### Run ethexporter
You'll need access to an ethereum Geth server to fetch balances. You can use [Infura.io](https://infura.io/setup) to quickly get an API key for a geth server.
`docker run -d -p 9015:9015 -e GETH="https://mainnet.infura.io/****KEYHERE" hunterlong/tokenexporter:latest`

## Pull from Dockerhub
Create a `addresses.txt` file with the correct format mentioned above.
```
docker run -d -v ~/tokenexporter:/app \
 -p 9021:9021 \
 -e GETH=https://mainnet.infura.io/****KEYHERE \
 hunterlong/tokenexporter:latest
```
The Docker image should be running with the default addresses.

## Prometheus Response
```
token_balance{name="etherdelta",symbol="$TEAK",contract="0x7dd7f56d697cc0f2b52bd55c057f378f1fe6ab4b",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="1ST",contract="0xAf30D2a7E90d7DC361c8C4585e9BB7D2F6f15bc7",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 385205.953499985641501145
token_balance{name="etherdelta",symbol="1WO",contract="0xfdbc1adc26f0f8f8606a5d63b7d3a3cd21c22b23",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 43811.35800826
token_balance{name="etherdelta",symbol="300",contract="0xaEc98A708810414878c3BCDF46Aad31dEd4a4557",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="A18",contract="0xBDe8f7820b5544a49D34F9dDeaCAbEDC7C0B5adc",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="ABT",contract="0xb98d4c97425d9908e66e53a6fdf673acca0be986",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 8041.990705949821978021
token_balance{name="etherdelta",symbol="ACC",contract="0x13f1b7fdfbe1fc66676d56483e21b1ecb40b58e2",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 6679.731723030207802662
token_balance{name="etherdelta",symbol="ADH",contract="0xE69a353b3152Dd7b706ff7dD40fe1d18b7802d31",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 39716.996
token_balance{name="etherdelta",symbol="ADI",contract="0x8810C63470d38639954c6B41AaC545848C46484a",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 6443.152228092911972885
token_balance{name="etherdelta",symbol="ADST",contract="0x422866a8F0b032c5cf1DfBDEf31A20F4509562b0",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 209622
token_balance{name="etherdelta",symbol="ADT",contract="0xD0D6D6C5Fe4a677D343cC433536BB717bAe167dD",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 1863273.947756648
token_balance{name="etherdelta",symbol="ADX",contract="0x4470BB87d77b963A013DB939BE332f927f2b992e",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 194192.1446
token_balance{name="etherdelta",symbol="AE",contract="0x5ca9a71b1d01849c0a95490cc00559717fcf0d1d",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 165041.401092921484358009
token_balance{name="etherdelta",symbol="AGI",contract="0x8eB24319393716668D768dCEC29356ae9CfFe285",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 1042938.86841104
token_balance{name="etherdelta",symbol="AION",contract="0x4CEdA7906a5Ed2179785Cd3A40A69ee8bc99C466",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 1192007.17387723
token_balance{name="etherdelta",symbol="AIR",contract="0x27dce1ec4d3f72c3e457cc50354f1f975ddef488",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 23765405.40501427
token_balance{name="etherdelta",symbol="AIX",contract="0x1063ce524265d5a3A624f4914acd573dD89ce988",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 38550.572871345222088238
token_balance{name="etherdelta",symbol="ALCO",contract="0x181a63746d3adcf356cbc73ace22832ffbb1ee5a",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 72050.07610484
token_balance{name="etherdelta",symbol="ALIS",contract="0xEA610B1153477720748DC13ED378003941d84fAB",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 129959.75341811347368985
token_balance{name="etherdelta",symbol="ALTS",contract="0x638ac149ea8ef9a1286c41b977017aa7359e6cfa",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 7398.9999995015
token_balance{name="etherdelta",symbol="AMB",contract="0x4dc3643dbc642b72c158e7f3d2ff232df61cb6ce",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 610026.146764510537824764
token_balance{name="etherdelta",symbol="AMIS",contract="0x949bed886c739f1a3273629b3320db0c5024c719",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 986810.027886806
token_balance{name="etherdelta",symbol="AMN",contract="0x737f98ac8ca59f2c68ad658e3c3d8c8963e40a4c",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="AMTC",contract="0x84936cF7630AA3e27Dd9AfF968b140d5AEE49F5a",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="ANT",contract="0x960b236A07cf122663c4303350609A66A7B288C0",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 122371.432648182401916902
token_balance{name="etherdelta",symbol="APIS",contract="0x4c0fbe1bb46612915e7967d2c3213cd4d87257ad",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 10429818.0
token_balance{name="etherdelta",symbol="APPC",contract="0x1a7a8bd9106f2b8d977e08582dc7d24c723ab0db",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 6663.126126585646472088
token_balance{name="etherdelta",symbol="APT",contract="0x23ae3c5b39b12f0693e05435eeaa1e51d8c61530",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="ARC",contract="0xAc709FcB44a43c35F0DA4e3163b117A17F3770f5",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 11410.833040644101512248
token_balance{name="etherdelta",symbol="ARCT",contract="0x1245ef80f4d9e02ed9425375e8f649b9221b31d8",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 2928535.34700729
token_balance{name="etherdelta",symbol="ARD",contract="0x75aa7b0d02532f3833b66c7f0ad35376d373ddf8",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 101617.068402650687284523
token_balance{name="etherdelta",symbol="ARN",contract="0xBA5F11b16B155792Cf3B2E6880E8706859A8AEB6",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 54612.83830791
token_balance{name="etherdelta",symbol="ART",contract="0xfec0cF7fE078a500abf15F1284958F22049c2C7e",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 129599.55440549934486142
token_balance{name="etherdelta",symbol="ARX",contract="0x7705FaA34B16EB6d77Dfc7812be2367ba6B0248e",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="ARXT",contract="0xb0D926c1BC3d78064F3e1075D5bD9A24F35Ae6C5",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 3210.44
token_balance{name="etherdelta",symbol="AST",contract="0x27054b13b1B798B345b591a4d22e6562d47eA75a",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 611472.9105
token_balance{name="etherdelta",symbol="ATH",contract="0x17052d51E954592C1046320c2371AbaB6C73Ef10",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="ATL",contract="0x78B7FADA55A64dD895D8c8c35779DD8b67fA8a05",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 106272.397990693048990241
token_balance{name="etherdelta",symbol="ATT",contract="0x887834d3b8d450b6bab109c252df3da286d73ce4",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 3275.944406670461196463
token_balance{name="etherdelta",symbol="ATTN",contract="0x6339784d9478da43106a429196772a029c2f177d",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 379000.0
token_balance{name="etherdelta",symbol="AVA",contract="0xeD247980396B10169BB1d36f6e278eD16700a60f",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="AVT",contract="0x0d88ed6e74bbfd96b831231638b66c05571e824f",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 53928.665305345456651211
token_balance{name="etherdelta",symbol="AX1",contract="0xcd4b4b0f3284a33ac49c67961ec6e111708318cf",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="AXP",contract="0x9af2c6B1A28D3d6BC084bd267F70e90d49741D5B",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 738147.83233755
token_balance{name="etherdelta",symbol="BANX",contract="0xf87f0d9153fea549c728ad61cb801595a68b73de",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="BAT",contract="0x0D8775F648430679A709E98d2b0Cb6250d2887EF",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 453204.253568495204036201
token_balance{name="etherdelta",symbol="BCBC",contract="0x7367a68039d4704f30bfbf6d948020c3b07dfc59",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 7350.0
token_balance{name="etherdelta",symbol="BCDN",contract="0x1e797Ce986C3CFF4472F7D38d5C4aba55DfEFE40",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="BCDT",contract="0xacfa209fb73bf3dd5bbfb1101b9bc999c49062a5",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 396177.325560001046242217
token_balance{name="etherdelta",symbol="BCL",contract="0xbc1234552EBea32B5121190356bBa6D3Bb225bb5",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="BCPT",contract="0x1c4481750daa5Ff521A2a7490d9981eD46465Dbd",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 1247971.760902492224110561
token_balance{name="etherdelta",symbol="BCV",contract="0x1014613e2b3cbc4d575054d4982e580d9b99d7b1",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="BDG",contract="0x1961B3331969eD52770751fC718ef530838b6dEE",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 2733643.454807774034363694
token_balance{name="etherdelta",symbol="BEE",contract="0x4D8fc1453a0F359e99c9675954e656D80d996FbF",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 306487.646261868153583527
token_balance{name="etherdelta",symbol="BERRY",contract="0x6aeb95f06cda84ca345c2de0f3b7f96923a44f4c",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 72284.70087315537927
token_balance{name="etherdelta",symbol="BET",contract="0x8aA33A7899FCC8eA5fBe6A608A109c3893A1B8b2",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 395745.327158375373709762
token_balance{name="etherdelta",symbol="BETR",contract="0x763186eb8d4856d536ed4478302971214febc6a9",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 108024.974755097942916499
token_balance{name="etherdelta",symbol="BKB",contract="0xb2bfeb70b903f1baac7f2ba2c62934c7e5b974c4",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 228361.52187341
token_balance{name="etherdelta",symbol="BKX",contract="0x45245bc59219eeaaf6cd3f382e078a461ff9de7b",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 118244.724914832203222873
token_balance{name="etherdelta",symbol="BLT",contract="0x107c4504cd79C5d2696Ea0030a8dD4e92601B82e",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 125702.356141089407691471
token_balance{name="etherdelta",symbol="BLX (Bullion)",contract="0xce59d29b09aae565feeef8e52f47c3cd5368c663",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 7617298.008200869120952936
token_balance{name="etherdelta",symbol="BLX (Iconomi)",contract="0xE5a7c12972f3bbFe70ed29521C8949b8Af6a0970",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 9540.32835465121192196
token_balance{name="etherdelta",symbol="BMC",contract="0xdf6ef343350780bf8c3410bf062e0c015b1dd671",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 127303.47265537
token_balance{name="etherdelta",symbol="BMT",contract="0xf028adee51533b1b47beaa890feb54a457f51e89",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 180632.014329067351834061
token_balance{name="etherdelta",symbol="BMX",contract="0x986EE2B944c42D017F52Af21c4c69B84DBeA35d8",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 183.155
token_balance{name="etherdelta",symbol="BNB",contract="0xb8c77482e45f1f44de1745f52c74426c631bdd52",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 102216.367700907897275724
token_balance{name="etherdelta",symbol="BNC",contract="0xdD6Bf56CA2ada24c683FAC50E37783e55B57AF9F",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="BNFT",contract="0xdA2C424Fc98c741c2d4ef2f42897CEfed897CA75",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="BNT",contract="0x1F573D6Fb3F13d689FF844B4cE37794d79a7FF1C",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 5508.770987653694405824
token_balance{name="etherdelta",symbol="BNTY",contract="0xd2d6158683aee4cc838067727209a0aaf4359de3",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 1169448.182879548824545749
token_balance{name="etherdelta",symbol="BON",contract="0xCc34366E3842cA1BD36c1f324d15257960fCC801",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 13742.680657096624169223
token_balance{name="etherdelta",symbol="BOP",contract="0x7f1e2c7d6a69bf34824d72c53b4550e895c0d8c2",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 51181.69383949
token_balance{name="etherdelta",symbol="BOU",contract="0xC2C63F23ec5E97efbD7565dF9Ec764FDc7d4e91d",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 3204.0
token_balance{name="etherdelta",symbol="BPT",contract="0x327682779bAB2BF4d1337e8974ab9dE8275A7Ca8",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 392764.680325486931705577
token_balance{name="etherdelta",symbol="BQX",contract="0x5Af2Be193a6ABCa9c8817001F45744777Db30756",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 60298.08272759
token_balance{name="etherdelta",symbol="BRAT",contract="0x9E77D5a1251b6F7D456722A6eaC6D2d5980bd891",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 9508888.07382649
token_balance{name="etherdelta",symbol="BSDC",contract="0xf26ef5e0545384b7dcc0f297f2674189586830df",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 3130151.335185518421209428
token_balance{name="etherdelta",symbol="BST",contract="0x509A38b7a1cC0dcd83Aa9d06214663D9eC7c7F4a",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="BTCE",contract="0x0886949c1b8C412860c4264Ceb8083d1365e86CF",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 9852289.29191807
token_balance{name="etherdelta",symbol="BTE",contract="0x73dd069c299a5d691e9836243bcaec9c8c1d8734",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 39161.38005166
token_balance{name="etherdelta",symbol="BTH",contract="0xfad572db566e5234ac9fc3d570c4edc0050eaa92",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 1892.231709147244257916
token_balance{name="etherdelta",symbol="BTHR",contract="0xa02e3bb9cebc03952601b3724b4940e0845bebcf",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="BTK",contract="0xdb8646f5b487b5dd979fac618350e85018f557d4",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 24966542.472282645253012133
token_balance{name="etherdelta",symbol="BTL (Battle)",contract="0x2accaB9cb7a48c3E82286F0b2f8798D201F4eC3f",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 10200.0
token_balance{name="etherdelta",symbol="BTL (Bitlle)",contract="0x92685E93956537c25Bb75D5d47fca4266dd628B8",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 14.0
token_balance{name="etherdelta",symbol="BTM",contract="0xcb97e65f07da24d46bcdd078ebebd7c6e6e3d750",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 172925.6382145
token_balance{name="etherdelta",symbol="BTQ",contract="0x16B0E62aC13a2fAeD36D18bce2356d25Ab3CfAD3",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="BTT",contract="0x080aa07e2c7185150d7e4da98838a8d2feac3dfc",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="BUC",contract="0xca3c18a65b802ec267f8f4802545e7f53d24c75e",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="BeerCoin",contract="0x74C1E4b8caE59269ec1D85D3D4F324396048F4ac",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="C20",contract="0x26E75307Fc0C021472fEb8F727839531F112f317",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 31088.851942929563199287
token_balance{name="etherdelta",symbol="CAG",contract="0x7d4b8Cce0591C9044a22ee543533b72E976E36C3",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 179195.187506945151085686
token_balance{name="etherdelta",symbol="CAN",contract="0x1d462414fe14cf489c7A21CaC78509f4bF8CD7c0",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 59823.614176
token_balance{name="etherdelta",symbol="CAR",contract="0x423e4322cdda29156b49a17dfbd2acc4b280600d",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="CARB",contract="0xa517a46baad6b054a76bd19c46844f717fe69fea",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="CAS",contract="0xe8780B48bdb05F928697A5e8155f672ED91462F7",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 36812.793284164213407659
token_balance{name="etherdelta",symbol="CAT (BitClave)",contract="0x1234567461d3f8db7496581774bd869c83d51c93",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 222578.803625692553151553
token_balance{name="etherdelta",symbol="CAT (Blockcat)",contract="0x56ba2Ee7890461f463F7be02aAC3099f6d5811A8",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 227216.370661747026109556
token_balance{name="etherdelta",symbol="CATs (BitClave)_Old",contract="0x68e14bb5A45B9681327E16E528084B9d962C1a39",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 101.0
token_balance{name="etherdelta",symbol="CC3",contract="0xc166038705FFBAb3794185b3a9D925632A1DF37D",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="CCC (CryptoCrashCourse)",contract="0x28577A6d31559bd265Ce3ADB62d0458550F7b8a7",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="CCC (ICONOMI)",contract="0xbe11eeb186e624b8f26a5045575a1340e4054552",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 788.0
token_balance{name="etherdelta",symbol="CCLC",contract="0xd348e07a2806505b856123045d27aeed90924b50",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="CCS",contract="0x315ce59fafd3a8d562b7ec1c8542382d2710b06c",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="CDL",contract="0x8a95ca448A52C0ADf0054bB3402dC5e09CD6B232",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="CDT",contract="0x177d39AC676ED1C67A2b268AD7F1E58826E5B0af",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 1106301.120223705660534409
token_balance{name="etherdelta",symbol="CDX",contract="0x6fFF3806Bbac52A20e0d79BC538d527f6a22c96b",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 589221.034680482128995966
token_balance{name="etherdelta",symbol="CFI",contract="0x12FEF5e57bF45873Cd9B62E9DBd7BFb99e32D73e",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 228719.916682190109440112
token_balance{name="etherdelta",symbol="CHSB",contract="0xba9d4199fab4f26efe3551d490e3821486f135ba",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 107903.02880751
token_balance{name="etherdelta",symbol="CK",contract="0x06012c8cf97bead5deae237070f9587f8e7a266d",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="CLL",contract="0x3dc9a42fa7afe57be03c58fd7f4411b1e466c508",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 31007.867209351011863223
token_balance{name="etherdelta",symbol="CLN",contract="0x4162178B78D6985480A308B2190EE5517460406D",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 200148.044554256299956188
token_balance{name="etherdelta",symbol="CLP",contract="0x7fce2856899a6806eeef70807985fc7554c66340",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="CMBT",contract="0x3edd235c3e840c1f29286b2e39370a255c7b6fdb",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="CMC",contract="0x7e667525521cF61352e2E01b50FaaaE7Df39749a",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="CMT",contract="0xf85fEea2FdD81d51177F6b8F35F0e6734Ce45F5F",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 84208.740261857350031026
token_balance{name="etherdelta",symbol="CNB",contract="0xEBf2F9E8De960f64ec0fDCDa6Cb282423133347B",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="CND",contract="0xd4c435f5b09f855c3317c8524cb1f586e42795fa",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 6743988.451394539489920198
token_balance{name="etherdelta",symbol="CO2",contract="0xB4b1D2C217EC0776584CE08D3DD98F90EDedA44b",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 1701.861837548996112052
token_balance{name="etherdelta",symbol="COB",contract="0xb2f7eb1f2c37645be61d73953035360e768d81e6",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 2106677.944591201472505514
token_balance{name="etherdelta",symbol="COFI",contract="0x3136eF851592aCf49CA4C825131E364170FA32b3",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 1216677.128472710193218083
token_balance{name="etherdelta",symbol="COIL",contract="0x0c91b015aba6f7b4738dcd36e7410138b29adc29",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="COSS",contract="0x65292eeadf1426cd2df1c4793a3d7519f253913b",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 2211888.7535987737784171
token_balance{name="etherdelta",symbol="COSS",contract="0x9e96604445ec19ffed9a5e8dd7b50a29c899a10c",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 696261.318425342420348844
token_balance{name="etherdelta",symbol="CPY",contract="0xf44745fbd41f6a1ba151df190db0564c5fcc4410",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 8591.239458427867558931
token_balance{name="etherdelta",symbol="CR7",contract="0x7f585b9130c64e9e9f470b618a7badd03d79ca7e",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 56898.12808393978666135
token_balance{name="etherdelta",symbol="CRB",contract="0xAef38fBFBF932D1AeF3B808Bc8fBd8Cd8E1f8BC5",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="CRED",contract="0x672a1AD4f667FB18A333Af13667aa0Af1F5b5bDD",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 1427350.517130627062783904
token_balance{name="etherdelta",symbol="CREDO",contract="0x4e0603e2a27a30480e5e3a4fe548e29ef12f64be",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 4889742.116858633681547715
token_balance{name="etherdelta",symbol="CRPT",contract="0x80a7e048f37a50500351c204cb407766fa3bae7f",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 102499.553475350475738324
token_balance{name="etherdelta",symbol="CRT",contract="0xF0da1186a4977226b9135d0613ee72e229EC3F4d",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 10000.0
token_balance{name="etherdelta",symbol="CTF",contract="0x4545750F39aF6Be4F237B6869D4EccA928Fd5A85",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 200.0003401455487988
token_balance{name="etherdelta",symbol="CTL",contract="0xbf4cfd7d1edeeea5f6600827411b41a21eb08abd",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="CTT",contract="0xE3Fa177AcecfB86721Cf6f9f4206bd3Bd672D7d5",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 343170.194636249654583082
token_balance{name="etherdelta",symbol="CTX",contract="0x662aBcAd0b7f345AB7FfB1b1fbb9Df7894f18e66",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 913900.747514647988166908
token_balance{name="etherdelta",symbol="CVC",contract="0x41e5560054824ea6b0732e656e3ad64e20e94e45",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 262602.71802231
token_balance{name="etherdelta",symbol="CXC",contract="0x2134057c0b461f898d375cead652acae62b59541",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="CXO",contract="0xb6EE9668771a79be7967ee29a63D4184F8097143",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 780378.010435622093347315
token_balance{name="etherdelta",symbol="CryptoCarbon",contract="0xE4c94d45f7Aef7018a5D66f44aF780ec6023378e",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="DAB",contract="0xdab0C31BF34C897Fb0Fe90D12EC9401caf5c36Ec",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="DADI",contract="0xfb2f26f266fb2805a387230f2aa0a331b4d96fba",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 48165.536667997425336249
token_balance{name="etherdelta",symbol="DAI",contract="0x89d24A6b4CcB1B6fAA2625fE562bDD9a23260359",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 10142.335504585539749786
token_balance{name="etherdelta",symbol="DALC",contract="0x07d9e49ea402194bf48a8276dafb16e4ed633317",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 1838.77706593
token_balance{name="etherdelta",symbol="DAN",contract="0x9B70740e708a083C6fF38Df52297020f5DfAa5EE",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 17029.0000430754
token_balance{name="etherdelta",symbol="DAO",contract="0xBB9bc244D798123fDe783fCc1C72d3Bb8C189413",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 3.0999999999999999
token_balance{name="etherdelta",symbol="DAT",contract="0x81c9151de0c8bafcd325a57e3db5a5df1cebf79c",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 1506929.556642474702290273
token_balance{name="etherdelta",symbol="DATABroker",contract="0x1b5f21ee98eed48d292e8e2d3ed82b40a9728a22",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="DATACoin",contract="0x0cf0ee63788a0849fe5297f3407f701e122cc023",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 991656.732236072755263194
token_balance{name="etherdelta",symbol="DAXT",contract="0x61725f3db4004afe014745b21dab1e1677cc328b",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="DCA",contract="0x386Faa4703a34a7Fdb19Bec2e14Fd427C9638416",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 50937386.283293147396518521
token_balance{name="etherdelta",symbol="DCL",contract="0x399A0e6FbEb3d74c85357439f4c8AeD9678a5cbF",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 1053.307
token_balance{name="etherdelta",symbol="DCN",contract="0x08d32b0da63e2C3bcF8019c9c5d849d7a9d791e6",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 134978712
token_balance{name="etherdelta",symbol="DDF",contract="0xcC4eF9EEAF656aC1a2Ab886743E98e97E090ed38",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 17805.015851304127039041
token_balance{name="etherdelta",symbol="DEB",contract="0x151202C9c18e495656f372281F493EB7698961D5",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="DENT",contract="0x3597bfd533a99c9aa083587b074434e61eb0a258",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 65371806.36474865
token_balance{name="etherdelta",symbol="DEPO",contract="0x7cF271966F36343Bf0150F25E5364f7961c58201",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="DGD",contract="0xE0B7927c4aF23765Cb51314A0E0521A9645F0E2A",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 974.881581947
token_balance{name="etherdelta",symbol="DGPT",contract="0xf6cFe53d6FEbaEEA051f400ff5fc14F0cBBDacA1",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 24631.115770900134237602
token_balance{name="etherdelta",symbol="DGX",contract="0x55b9a11c2e8351b4Ffc7b11561148bfaC9977855",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="DICE",contract="0x2e071D2966Aa7D8dECB1005885bA1977D6038A65",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 72342.3664246006689712
token_balance{name="etherdelta",symbol="DIVX",contract="0x13f11C9905A08ca76e3e853bE63D4f0944326C72",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 2318.169277951454367244
token_balance{name="etherdelta",symbol="DLT",contract="0x07e3c70653548b04f0a75970c1f81b4cbbfb606f",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 263772.6388489228004388
token_balance{name="etherdelta",symbol="DMT",contract="0x2ccbFF3A042c68716Ed2a2Cb0c544A9f1d1935E1",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 104589.05495575
token_balance{name="etherdelta",symbol="DNT",contract="0x0abdace70d3790235af448c88547603b945604ea",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 2232968.880693111729999642
token_balance{name="etherdelta",symbol="DNX",contract="0xE43E2041dc3786e166961eD9484a5539033d10fB",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="DOW",contract="0x76974c7b79dc8a6a109fd71fd7ceb9e40eff5382",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="DPP",contract="0x01b3Ec4aAe1B8729529BEB4965F27d008788B0EB",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 9478.917231283024996667
token_balance{name="etherdelta",symbol="DRGN",contract="0x419c4db4b9e25d6db2ad9691ccb832c8d9fda05e",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 1758374.085893228207308057
token_balance{name="etherdelta",symbol="DROP (dropil)",contract="0x4672bad527107471cb5067a887f4656d585a8a31",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 273009.696094634113714938
token_balance{name="etherdelta",symbol="DROP (droplex)",contract="0x3c75226555FC496168d48B88DF83B95F16771F37",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 113321
token_balance{name="etherdelta",symbol="DRP",contract="0x621d78f2ef2fd937bfca696cabaf9a779f59b3ed",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 106274.05
token_balance{name="etherdelta",symbol="DRP (Dripcoin)",contract="0x2799d90c6d44cb9aa5fbc377177f16c33e056b82",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="DSC",contract="0x1e09BD8Cadb441632e441Db3e1D79909EE0A2256",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="DTR",contract="0xd234bf2410a0009df9c3c63b610c09738f18ccd7",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 98378.33274505
token_balance{name="etherdelta",symbol="DTX",contract="0x765f0c16d1ddc279295c1a7c24b0883f62d33f75",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="DUBI",contract="0xd4cffeef10f60eca581b5e1146b5aca4194a4c3b",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 55.586066415723106015
token_balance{name="etherdelta",symbol="Devcon2 Token",contract="0xdd94De9cFE063577051A5eb7465D08317d8808B6",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="EAGLE",contract="0x994f0dffdbae0bbf09b652d6f11a493fd33f42b9",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 188326.7779147367247572
token_balance{name="etherdelta",symbol="ECN",contract="0xa578acc0cb7875781b7880903f4594d13cfa8b98",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 67838.26
token_balance{name="etherdelta",symbol="ECO2",contract="0x17F93475d2A978f527c3f7c44aBf44AdfBa60D5C",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 68218.87
token_balance{name="etherdelta",symbol="EDC",contract="0xfa1de2ee97e4c10c94c91cb2b5062b89fb140b82",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="EDG",contract="0x08711D3B02C8758F2FB3ab4e80228418a7F8e39c",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 33067
token_balance{name="etherdelta",symbol="EDO",contract="0xced4e93198734ddaff8492d525bd258d49eb388e",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 32877.789010511122109913
token_balance{name="etherdelta",symbol="LEDU",contract="0x5b26C5D0772E5bbaC8b3182AE9a13f9BB2D03765",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 116307.1702857
token_balance{name="etherdelta",symbol="EHT",contract="0xf9F0FC7167c311Dd2F1e21E9204F87EBA9012fB2",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="ELF",contract="0xbf2179859fc6d5bee9bf9158632dc51678a4100e",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 589.913798917703294004
token_balance{name="etherdelta",symbol="ELIX",contract="0xc8C6A31A4A806d3710A7B38b7B296D2fABCCDBA8",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 189013.17371808373622433
token_balance{name="etherdelta",symbol="ELTCOIN",contract="0x44197a4c44d6a059297caf6be4f7e172bd56caaf",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 1847639.0904806
token_balance{name="etherdelta",symbol="EMON",contract="0xb67b88a25708a35ae7c2d736d398d268ce4f7f83",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 395.0
token_balance{name="etherdelta",symbol="EMONT",contract="0x95daaab98046846bf4b2853e23cba236fa394a31",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 4132.68036555
token_balance{name="etherdelta",symbol="EMT",contract="0x9501BFc48897DCEEadf73113EF635d2fF7ee4B97",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 304828.19468345665364502
token_balance{name="etherdelta",symbol="EMV",contract="0xB802b24E0637c2B87D2E8b7784C055BBE921011a",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 150630.17
token_balance{name="etherdelta",symbol="ENJ",contract="0xF629cBd94d3791C9250152BD8dfBDF380E2a3B9c",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 1774724.024828353964329669
token_balance{name="etherdelta",symbol="ENTRP",contract="0x5BC7e5f0Ab8b2E10D2D0a3F21739FCe62459aeF3",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 49039.261089662444454811
token_balance{name="etherdelta",symbol="EPX",contract="0x35BAA72038F127f9f8C8f9B491049f64f377914d",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="EOS",contract="0x86fa049857e0209aa7d9e616f7eb3b3b78ecfdb0",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 143155.949223553397128994
token_balance{name="etherdelta",symbol="ESZ",contract="0xe8a1df958be379045e2b46a31a98b93a2ecdfded",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 70267.961576840672454309
token_balance{name="etherdelta",symbol="ETBS",contract="0x1b9743f556d65e757c4c650b4555baf354cb8bd3",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 15164.261709723576
token_balance{name="etherdelta",symbol="ETHB",contract="0x3a26746Ddb79B1B8e4450e3F4FFE3285A307387E",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 206228.86356683
token_balance{name="etherdelta",symbol="ETCH",contract="0xdd74a7a3769fa72561b3a69e65968f49748c690c",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 21011.89729001129111045
token_balance{name="etherdelta",symbol="ETR",contract="0x6927C69fb4daf2043fbB1Cb7b86c5661416bea29",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0.44
token_balance{name="etherdelta",symbol="EURT",contract="0xabdf147870235fcfc34153828c769a70b3fae01f",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="EVE",contract="0x923108a439C4e8C2315c4f6521E5cE95B44e9B4c",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 2136028.685870220494213388
token_balance{name="etherdelta",symbol="EVN",contract="0xd780Ae2Bf04cD96E577D3D014762f831d97129d0",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 170085.894488443952831351
token_balance{name="etherdelta",symbol="EVX",contract="0xf3db5fa2c66b7af3eb0c0b782510816cbe4813b8",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 47958.983
token_balance{name="etherdelta",symbol="EXMR",contract="0xc98e0639c6d2ec037a615341c369666b110e80e5",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 744643.12723013
token_balance{name="etherdelta",symbol="E₹",contract="0xb67734521eAbBE9C773729dB73E16CC2dfb20A58",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 10009110000.0
token_balance{name="etherdelta",symbol="FAM",contract="0x190e569bE071F40c704e15825F285481CB74B6cC",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 251.689117647056
token_balance{name="etherdelta",symbol="FKX",contract="0x009e864923b49263c7F10D19B7f8Ab7a9A5AAd33",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 413177.265362988350659704
token_balance{name="etherdelta",symbol="FLIXX",contract="0xf04a8ac553FceDB5BA99A64799155826C136b0Be",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 1148374.051475376218207151
token_balance{name="etherdelta",symbol="FLMC",contract="0x04cC783b450b8D11F3C7d00DD03fDF7FB51fE9F2",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 2000000.0
token_balance{name="etherdelta",symbol="FLP",contract="0x3a1Bda28AdB5B0a812a7CF10A1950c920F79BcD3",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 209359.645393582432443069
token_balance{name="etherdelta",symbol="FLUZ",contract="0x954b5de09a55e59755acbda29e1eb74a45d30175",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 49182.97559220045351284
token_balance{name="etherdelta",symbol="FLX",contract="0x70b147e01e9285e7ce68b9ba437fe3a9190e756a",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 2341065.98863230695424394
token_balance{name="etherdelta",symbol="FND",contract="0x4df47b4969b2911c966506e3592c41389493953b",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 23348.551195793921246777
token_balance{name="etherdelta",symbol="FRD",contract="0x0ABeFb7611Cb3A01EA3FaD85f33C3C934F8e2cF4",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 1447252.553938369812752964
token_balance{name="etherdelta",symbol="FTC",contract="0xe6f74dcfa0e20883008d8c16b6d9a329189d0c30",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 2200.0
token_balance{name="etherdelta",symbol="FTR",contract="0x2023DCf7c438c8C8C0B0F28dBaE15520B4f3Ee20",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 85275.87697
token_balance{name="etherdelta",symbol="FUCK",contract="0x65be44c747988fbf606207698c944df4442efe19",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 122846.128
token_balance{name="etherdelta",symbol="FUEL",contract="0xEA38eAa3C86c8F9B751533Ba2E562deb9acDED40",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 2292941.931887991135878505
token_balance{name="etherdelta",symbol="FUN",contract="0x419D0d8BdD9aF5e606Ae2232ed285Aff190E711b",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 6894549.67372932
token_balance{name="etherdelta",symbol="FYN",contract="0x88FCFBc22C6d3dBaa25aF478C578978339BDe77a",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 4433.733745642151697904
token_balance{name="etherdelta",symbol="GAM",contract="0xf67451dc8421f0e0afeb52faa8101034ed081ed9",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 1
token_balance{name="etherdelta",symbol="GAVEL",contract="0x708876f486e448ee89eb332bfbc8e593553058b9",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="GBT",contract="0x7585F835ae2d522722d2684323a0ba83401f32f5",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 10128890225462329
token_balance{name="etherdelta",symbol="GBX",contract="0x12fCd6463E66974cF7bBC24FFC4d40d6bE458283",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 225047.028
token_balance{name="etherdelta",symbol="GCP",contract="0xdb0F69306FF8F949f258E83f6b87ee5D052d0b23",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 1289.9124201
token_balance{name="etherdelta",symbol="GEE",contract="0x4F4f0Db4de903B88f2B1a2847971E231D54F8fd3",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 83.333
token_balance{name="etherdelta",symbol="GELD",contract="0x24083bb30072643c3bb90b44b7285860a755e687",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 495018.95
token_balance{name="etherdelta",symbol="GIM",contract="0xaE4f56F072c34C0a65B3ae3E4DB797D831439D93",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 6855297.53648772
token_balance{name="etherdelta",symbol="GMT",contract="0xb3Bd49E28f8F832b8d1E246106991e546c323502",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 12342446.021788205261956559
token_balance{name="etherdelta",symbol="GNO",contract="0x6810e776880C02933D47DB1b9fc05908e5386b96",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 67.550646483380799665
token_balance{name="etherdelta",symbol="GNT",contract="0xa74476443119A942dE498590Fe1f2454d7D4aC0d",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 52.60000000000000011
token_balance{name="etherdelta",symbol="GOLDX",contract="0xeAb43193CF0623073Ca89DB9B712796356FA7414",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 2236983747668011
token_balance{name="etherdelta",symbol="GRID",contract="0x12b19d3e2ccc14da04fae33e63652ce469b3f2fd",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 884698.541645127303
token_balance{name="etherdelta",symbol="GTC",contract="0xB70835D7822eBB9426B56543E391846C107bd32C",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="GTKT",contract="0x025abad9e518516fdaafbdcdb9701b37fb7ef0fa",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 2
token_balance{name="etherdelta",symbol="GUP",contract="0xf7B098298f7C69Fc14610bf71d5e02c60792894C",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 11202.646
token_balance{name="etherdelta",symbol="GVT",contract="0x103c3A209da59d3E7C4A89307e66521e081CFDF0",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 10155.38114902931030881
token_balance{name="etherdelta",symbol="GXC",contract="0x58ca3065c0f24c7c96aee8d6056b5b5decf9c2f8",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 576088.3772322022
token_balance{name="etherdelta",symbol="GXVC",contract="0x22F0AF8D78851b72EE799e05F54A77001586B18A",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="GZE",contract="0x8C65e992297d5f092A756dEf24F4781a280198Ff",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 266459.211794511300762262
token_balance{name="etherdelta",symbol="HAT",contract="0x9002D4485b7594e3E850F0a206713B305113f69e",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 556.413602866596
token_balance{name="etherdelta",symbol="HDG",contract="0xffe8196bc259e8dedc544d935786aa4709ec3e64",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 81751.349512176560650455
token_balance{name="etherdelta",symbol="HGT",contract="0xba2184520A1cC49a6159c57e61E1844E085615B6",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 2686618.72435303
token_balance{name="etherdelta",symbol="HIG",contract="0xa9240fBCAC1F0b9A6aDfB04a53c8E3B0cC1D1444",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="HKG",contract="0x14F37B574242D366558dB61f3335289a5035c506",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 93827.068
token_balance{name="etherdelta",symbol="HKY",contract="0x88ac94d5d175130347fc95e109d77ac09dbf5ab7",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 90000.0
token_balance{name="etherdelta",symbol="HMQ",contract="0xcbCC0F036ED4788F63FC0fEE32873d6A7487b908",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 173651.92632952
token_balance{name="etherdelta",symbol="HODL",contract="0xb45d7Bc4cEBcAB98aD09BABDF8C818B2292B672c",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="HORSE",contract="0x5B0751713b2527d7f002c0c4e2a37e1219610A6B",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 359983.565256527870741113
token_balance{name="etherdelta",symbol="HOT (HoloToken)",contract="0x6c6ee5e31d828de241282b9606c8e98ea48526e2",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="HOT (Hydro Protocol)",contract="0x9af839687f6c94542ac5ece2e317daae355493a1",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 895.019
token_balance{name="etherdelta",symbol="HST",contract="0x554C20B7c486beeE439277b4540A434566dC4C02",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 196967.978105711730527435
token_balance{name="etherdelta",symbol="HVN",contract="0xC0Eb85285d83217CD7c891702bcbC0FC401E2D9D",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 1490414.86118352
token_balance{name="etherdelta",symbol="Hdp",contract="0xe9ff07809ccff05dae74990e25831d0bc5cbe575",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 1000000.0
token_balance{name="etherdelta",symbol="Hdp.ф",contract="0x84543f868ec1b1fac510d49d13c069f64cd2d5f9",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 1000000.0
token_balance{name="etherdelta",symbol="ICE",contract="0x5a84969bb663fb64F6d015DcF9F622Aedc796750",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 70822.861000010474248939
token_balance{name="etherdelta",symbol="ICN",contract="0x888666CA69E0f178DED6D75b5726Cee99A87D698",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 320803.470012820972493524
token_balance{name="etherdelta",symbol="ICO",contract="0xa33e729bf4fdeb868b534e1f20523463d9c46bee",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0.168
token_balance{name="etherdelta",symbol="ICOS",contract="0x014B50466590340D41307Cc54DCee990c8D58aa8",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 1033.447071
token_balance{name="etherdelta",symbol="ICX",contract="0xb5a5f22694352c15b00323844ad545abb2b11028",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 3399.363935594905539432
token_balance{name="etherdelta",symbol="IDEA",contract="0x814cafd4782d2e728170fda68257983f03321c58",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 5272481
token_balance{name="etherdelta",symbol="IFT",contract="0x7654915a1b82d6d2d0afc37c52af556ea8983c7e",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 616709.252947046071496838
token_balance{name="etherdelta",symbol="IIC",contract="0x16662f73df3e79e54c6c5938b4313f92c524c120",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 69.0
token_balance{name="etherdelta",symbol="IKB",contract="0x88AE96845e157558ef59e9Ff90E766E22E480390",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="IMC",contract="0xe3831c5A982B279A198456D577cfb90424cb6340",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="IMT",contract="0x22E5F62D0FA19974749faa194e3d3eF6d89c08d7",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="IND",contract="0xf8e386EDa857484f5a12e4B5DAa9984E06E73705",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 350184.502347055781056132
token_balance{name="etherdelta",symbol="INS",contract="0x5b2e4a700dfbc560061e957edec8f6eeeb74a320",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 15669.9460119471
token_balance{name="etherdelta",symbol="INSTAR",contract="0xc72fe8e3dd5bef0f9f31f259399f301272ef2a2d",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 158477.763451549433845775
token_balance{name="etherdelta",symbol="INXT",contract="0xa8006c4ca56f24d6836727d106349320db7fef82",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 5062.40002652
token_balance{name="etherdelta",symbol="IPL",contract="0x64CdF819d3E75Ac8eC217B3496d7cE167Be42e80",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 279102.503885205487655356
token_balance{name="etherdelta",symbol="IPSX",contract="0x001f0aa5da15585e5b2305dbab2bac425ea71007",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 114820.23137782873093411
token_balance{name="etherdelta",symbol="ITC",contract="0x5e6b6d9abad9093fdc861ea1600eba1b355cd940",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 50.626707207656265578
token_balance{name="etherdelta",symbol="ITT",contract="0x0aeF06DcCCC531e581f0440059E6FfCC206039EE",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 344519.17976923
token_balance{name="etherdelta",symbol="IXT",contract="0xfca47962d45adfdfd1ab2d972315db4ce7ccf094",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 481049.12694948
token_balance{name="etherdelta",symbol="IoT",contract="0xc34b21f6f8e51cc965c2393b3ccfa3b82beb2403",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="J8T",contract="0x0d262e5dc4a06a0f1c90ce79c7a60c09dfc884e4",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 742462.98847233
token_balance{name="etherdelta",symbol="JBX",contract="0x0Aaf561eFF5BD9c8F911616933F84166A17cfE0C",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="JET",contract="0x8727c112c712c4a03371ac87a74dd6ab104af768",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 76462.62619417997157927
token_balance{name="etherdelta",symbol="JNT",contract="0xa5Fd1A791C4dfcaacC963D4F73c6Ae5824149eA7",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 20039.80828116175488377
token_balance{name="etherdelta",symbol="JetCoins",contract="0x773450335eD4ec3DB45aF74f34F2c85348645D39",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="KC",contract="0x0D6DD9f68d24EC1d5fE2174f3EC8DAB52B52BaF5",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 100.0
token_balance{name="etherdelta",symbol="KEE",contract="0x72D32ac1c5E66BfC5b08806271f8eEF915545164",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 13
token_balance{name="etherdelta",symbol="KEY (SelfKey)",contract="0x4CC19356f2D37338b9802aa8E8fc58B0373296E7",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 3258373.578126201543504822
token_balance{name="etherdelta",symbol="KEY (BihuKey)",contract="0x4cd988afbad37289baaf53c13e98e2bd46aaea8c",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 9534048.206078418361581265
token_balance{name="etherdelta",symbol="KICK",contract="0x27695E09149AdC738A978e9A678F99E4c39e9eb9",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 4357052.62567285
token_balance{name="etherdelta",symbol="KIN",contract="0x818Fc6C2Ec5986bc6E2CBf00939d90556aB12ce5",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 11725956244.010848250104429055
token_balance{name="etherdelta",symbol="KNC",contract="0xdd974D5C2e2928deA5F71b9825b8b646686BD200",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 288185.606018396798594892
token_balance{name="etherdelta",symbol="KPR",contract="0xb5c33f965c8899d255c34cdd2a3efa8abcbb3dea",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="KZN",contract="0x9541FD8B9b5FA97381783783CeBF2F5fA793C262",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 118042.18809272
token_balance{name="etherdelta",symbol="LA",contract="0xE50365f5D679CB98a1dd62D6F6e58e59321BcdDf",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 216877.547158548737033518
token_balance{name="etherdelta",symbol="LALA",contract="0xfD107B473AB90e8Fbd89872144a3DC92C40Fa8C9",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 208221.825631236056116051
token_balance{name="etherdelta",symbol="LDC",contract="0x5102791ca02fc3595398400bfe0e33d7b6c82267",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 4209.64105682468439559
token_balance{name="etherdelta",symbol="LEMO",contract="0xd6e354F07319e2474491D8c7c712137bEe6862a2",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="LFR",contract="0xc798cd1c49db0e297312e4c682752668ce1db2ad",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="LGR",contract="0x2eb86e8fc520e0f6bb5d9af08f924fe70558ab89",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 3463150.46311109
token_balance{name="etherdelta",symbol="LIF",contract="0xEB9951021698B42e4399f9cBb6267Aa35F82D59D",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 500.0
token_balance{name="etherdelta",symbol="LIFE",contract="0xff18dbc487b4c2e3222d115952babfda8ba52f5f",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 49483370.695284705322112774
token_balance{name="etherdelta",symbol="LINK (Chainlink)",contract="0x514910771af9ca656af840dff83e8264ecf986ca",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 1659187.329393180652098503
token_balance{name="etherdelta",symbol="LINK Platform",contract="0xe2e6d4be086c6938b53b22144855eef674281639",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 975.519608489447373077
token_balance{name="etherdelta",symbol="LIVE",contract="0x24A77c1F17C547105E14813e517be06b0040aa76",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 249659.901292874445582911
token_balance{name="etherdelta",symbol="LNC",contract="0x63e634330A20150DbB61B15648bC73855d6CCF07",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 137060.195710572249400454
token_balance{name="etherdelta",symbol="LNC-Linker Coin",contract="0x6beb418fc6e1958204ac8baddcf109b8e9694966",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 386198.089067085813715643
token_balance{name="etherdelta",symbol="LND",contract="0x0947b0e6D821378805c9598291385CE7c791A6B2",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="LOC",contract="0x5e3346444010135322268a4630d2ed5f8d09446c",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 29865.189953374478768343
token_balance{name="etherdelta",symbol="LOCI",contract="0x9c23d67aea7b95d80942e3836bcdf7e708a747c2",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 38800.627095763540811719
token_balance{name="etherdelta",symbol="LOCUS",contract="0xC64500DD7B0f1794807e67802F8Abbf5F8Ffb054",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="LOK",contract="0x21ae23b882a340a22282162086bc98d3e2b73018",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="LOOM",contract="0xa4e8c3ec456107ea67d3075bf9e3df3a75823db0",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 44.00017640999996
token_balance{name="etherdelta",symbol="LRC",contract="0xEF68e7C694F40c8202821eDF525dE3782458639f",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 981474.635271884053527632
token_balance{name="etherdelta",symbol="LUCK",contract="0xFB12e3CcA983B9f59D90912Fd17F8D745A8B2953",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 2887888
token_balance{name="etherdelta",symbol="LUM",contract="0xa89b5934863447f6e4fc53b315a93e873bda69a3",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 38025.455263334120585514
token_balance{name="etherdelta",symbol="LUN",contract="0xfa05A73FfE78ef8f1a739473e462c54bae6567D9",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 1635.995826561469520474
token_balance{name="etherdelta",symbol="M-ETH",contract="0x3f4b726668da46f5e0e75aa5d478acec9f38210f",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="MAD",contract="0x5b09a0371c1da44a8e24d36bf5deb1141a84d875",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 86576.0
token_balance{name="etherdelta",symbol="MANA",contract="0x0F5D2fB29fb7d3CFeE444a200298f468908cC942",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 2003968.534138799355221056
token_balance{name="etherdelta",symbol="MBRS",contract="0x386467f1f3ddbe832448650418311a479eecfc57",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 668931
token_balance{name="etherdelta",symbol="MCAP",contract="0x93E682107d1E9defB0b5ee701C71707a4B2E46Bc",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 38759.86371456
token_balance{name="etherdelta",symbol="MCI",contract="0x138A8752093F4f9a79AaeDF48d4B9248fab93c9C",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 658982.171184157148286397
token_balance{name="etherdelta",symbol="MCO",contract="0xB63B606Ac810a52cCa15e44bB630fd42D8d1d83d",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 4197.17641541
token_balance{name="etherdelta",symbol="MDA",contract="0x51DB5Ad35C671a87207d88fC11d593AC0C8415bd",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 26904.310535575187385698
token_balance{name="etherdelta",symbol="MESH",contract="0x01f2acf2914860331c1cb1a9acecda7475e06af8",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="MEST",contract="0x5b8d43ffde4a2982b9a5387cdf21d54ead64ac8d",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="MGO",contract="0x40395044ac3c0c57051906da938b54bd6557f212",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 53711.45728142
token_balance{name="etherdelta",symbol="MIT",contract="0xe23cd160761f63FC3a1cF78Aa034b6cdF97d3E0C",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="MKR",contract="0x9f8F72aA9304c8B593d555F12eF6589cC3A579A2",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 10.44216219634276957
token_balance{name="etherdelta",symbol="MLN",contract="0xBEB9eF514a379B997e0798FDcC901Ee474B6D9A1",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 2032.218612625938838304
token_balance{name="etherdelta",symbol="MNE",contract="0x1a95B271B0535D15fa49932Daba31BA612b52946",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="MNT",contract="0xA9877b1e05D035899131DBd1e403825166D09f92",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="MNTP",contract="0x83cee9e086a77e492ee0bb93c2b0437ad6fdeccc",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 30161.469084042250294388
token_balance{name="etherdelta",symbol="MOD",contract="0x957c30aB0426e0C93CD8241E2c60392d08c6aC8e",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 9228
token_balance{name="etherdelta",symbol="MRP",contract="0x21f0F0fD3141Ee9E11B3d7f13a1028CD515f459c",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="MRV",contract="0xAB6CF87a50F17d7F5E1FEaf81B6fE9FfBe8EBF84",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 18955.81738196821421255
token_balance{name="etherdelta",symbol="MSP",contract="0x68AA3F232dA9bdC2343465545794ef3eEa5209BD",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 104863.656685320181295588
token_balance{name="etherdelta",symbol="MTH",contract="0xaF4DcE16Da2877f8c9e00544c93B62Ac40631F16",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 2252785.42288
token_balance{name="etherdelta",symbol="MTL",contract="0xF433089366899D83a9f26A773D59ec7eCF30355e",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 4384.66376948
token_balance{name="etherdelta",symbol="MTN",contract="0x41dbecc1cdc5517c6f76f6a6e836adbee2754de3",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 79015.849826390335876061
token_balance{name="etherdelta",symbol="MTR",contract="0x7FC408011165760eE31bE2BF20dAf450356692Af",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="MTRc",contract="0x1e49fF77c355A3e38D6651ce8404AF0E48c5395f",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 9169.006640797723
token_balance{name="etherdelta",symbol="MTX",contract="0x0AF44e2784637218dD1D32A322D44e603A8f0c6A",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 40370.540130318345832127
token_balance{name="etherdelta",symbol="MWAT",contract="0x6425c6be902d692ae2db752b3c268afadb099d3b",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 841508.460970721297510247
token_balance{name="etherdelta",symbol="MYD",contract="0xf7e983781609012307f2514f63D526D83D24F466",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="MYST",contract="0xa645264C5603E96c3b0B078cdab68733794B0A71",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 27519.4420625
token_balance{name="etherdelta",symbol="NAVI",contract="0x588047365df5ba589f923604aac23d673555c623",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 358068.6
token_balance{name="etherdelta",symbol="NBAI",contract="0x17f8aFB63DfcDcC90ebE6e84F060Cc306A98257D",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="NCT",contract="0x9e46a38f5daabe8683e10793b06749eef7d733d1",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 60142.893995513861481201
token_balance{name="etherdelta",symbol="NDC",contract="0xa54ddc7b3cce7fc8b1e3fa0256d0db80d2c10970",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 122944.53304738043551576
token_balance{name="etherdelta",symbol="NET",contract="0xcfb98637bcae43C13323EAa1731cED2B716962fD",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 23366.595237527416187301
token_balance{name="etherdelta",symbol="NEU",contract="0xa823e6722006afe99e91c30ff5295052fe6b8e32",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 6389.495406917049661918
token_balance{name="etherdelta",symbol="NGC",contract="0x72dd4b6bd852a3aa172be4d6c5a6dbec588cf131",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 2145.465592112919756329
token_balance{name="etherdelta",symbol="NIMFA",contract="0xe26517A9967299453d3F1B48Aa005E6127e67210",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 87505.062027042700110432
token_balance{name="etherdelta",symbol="NMR",contract="0x1776e1F26f98b1A5dF9cD347953a26dd3Cb46671",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 2250.637270112724041455
token_balance{name="etherdelta",symbol="NOX",contract="0xec46f8207d766012454c408de210bcbc2243e71c",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 3531.113143908390345543
token_balance{name="etherdelta",symbol="NPER",contract="0x4ce6b362bc77a24966dda9078f9cef81b3b886a7",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 31.868529432477965464
token_balance{name="etherdelta",symbol="NULS",contract="0xb91318f35bdb262e9423bc7c7c2a3a93dd93c92c",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 4522.925216441938522967
token_balance{name="etherdelta",symbol="NXX",contract="0x7627de4b93263a6a7570b8dafa64bae812e5c394",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 610085.19631565
token_balance{name="etherdelta",symbol="NXX OLD",contract="0x5c6183d10A00CD747a6Dbb5F658aD514383e9419",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 26648547.0683903
token_balance{name="etherdelta",symbol="NxC",contract="0x45e42D659D9f9466cD5DF622506033145a9b89Bc",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 940.213
token_balance{name="etherdelta",symbol="OAX",contract="0x701C244b988a513c945973dEFA05de933b23Fe1D",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 37189.714213822758348856
token_balance{name="etherdelta",symbol="OCC",contract="0x0235fe624e044a05eed7a43e16e3083bc8a4287a",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 2966410172.464279713158414236
token_balance{name="etherdelta",symbol="OHNI",contract="0x6f539a9456a5bcb6334a1a41207c3788f5825207",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="OLD_MKR",contract="0xc66ea802717bfb9833400264dd12c2bceaa34a6d",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 238.904353539690089205
token_balance{name="etherdelta",symbol="OMG",contract="0xd26114cd6EE289AccF82350c8d8487fedB8A0C07",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 68709.439044949743412472
token_balance{name="etherdelta",symbol="ONEK",contract="0xb23be73573bc7e03db6e5dfc62405368716d28a8",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 32.743347311026141668
token_balance{name="etherdelta",symbol="OPT",contract="0x4355fC160f74328f9b383dF2EC589bB3dFd82Ba0",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 4093576.581187685182465174
token_balance{name="etherdelta",symbol="OST",contract="0x2C4e8f2D746113d0696cE89B35F0d8bF88E0AEcA",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 56493.417971070704717959
token_balance{name="etherdelta",symbol="Ox Fina",contract="0x65a15014964f2102ff58647e16a16a6b9e14bcf6",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 8525.5
token_balance{name="etherdelta",symbol="PAL",contract="0xfeDAE5642668f8636A11987Ff386bfd215F942EE",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="PARETO",contract="0xea5f88e54d982cbb0c441cde4e79bc305e5b43bc",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 1255253.987316451717756306
token_balance{name="etherdelta",symbol="PATENTS",contract="0x694404595e3075a942397f466aacd462ff1a7bd0",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="PATH",contract="0xF813F3902bBc00A6DCe378634d3B79D84F9803d7",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="PAY",contract="0xB97048628DB6B661D4C2aA833e95Dbe1A905B280",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 190819.795962163681683774
token_balance{name="etherdelta",symbol="PBL",contract="0x55648de19836338549130b1af587f16bea46f66b",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 280324.31403597273532335
token_balance{name="etherdelta",symbol="PBT",contract="0xF4c07b1865bC326A3c01339492Ca7538FD038Cc0",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="PCH",contract="0xfcAC7A7515e9A9d7619fA77A1fa738111f66727e",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="PCL",contract="0x3618516F45CD3c913F81F9987AF41077932Bc40d",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 487671.58029318
token_balance{name="etherdelta",symbol="PCLOLD",contract="0x53148Bb4551707edF51a1e8d7A93698d18931225",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 391409.91691789
token_balance{name="etherdelta",symbol="PET",contract="0x5884969Ec0480556E11d119980136a4C17eDDEd1",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="PETRO",contract="0xec18f898b4076a3e18f1089d33376cc380bde61d",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 28360.2836
token_balance{name="etherdelta",symbol="PEXT",contract="0x55c2A0C171D920843560594dE3d6EEcC09eFc098",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 12554.9889
token_balance{name="etherdelta",symbol="PIPL",contract="0xE64509F0bf07ce2d29A7eF19A8A9bc065477C1B4",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 3900.0
token_balance{name="etherdelta",symbol="PIX",contract="0x8eFFd494eB698cc399AF6231fCcd39E08fd20B15",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 174203
token_balance{name="etherdelta",symbol="PLASMA",contract="0x59416A25628A76b4730eC51486114c32E0B582A1",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="PLAY",contract="0xE477292f1B3268687A29376116B0ED27A9c76170",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 519740.778987998040404273
token_balance{name="etherdelta",symbol="PLBT",contract="0x0AfFa06e7Fbe5bC9a764C979aA66E8256A631f02",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 5693.979338
token_balance{name="etherdelta",symbol="PLR",contract="0xe3818504c1B32bF1557b16C238B2E01Fd3149C17",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 3689428.275256761277417895
token_balance{name="etherdelta",symbol="PLU",contract="0xD8912C10681D8B21Fd3742244f44658dBA12264E",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 7411.247876086240272636
token_balance{name="etherdelta",symbol="POE",contract="0x0e0989b1f9b8a38983c2ba8053269ca62ec9b195",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 4071183.18002581
token_balance{name="etherdelta",symbol="POIN",contract="0x43f6a1be992dee408721748490772b15143ce0a7",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="POLY",contract="0x9992eC3cF6A55b00978cdDF2b27BC6882d88D1eC",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 404366.986977112096536496
token_balance{name="etherdelta",symbol="POOL",contract="0x779B7b713C86e3E6774f5040D9cCC2D43ad375F8",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 4991.4236878
token_balance{name="etherdelta",symbol="POS",contract="0xee609fe292128cad03b786dbb9bc2634ccdbe7fc",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 44415.326376962744855487
token_balance{name="etherdelta",symbol="POWR",contract="0x595832f8fc6bf59c85c527fec3740a1b7a361269",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 257030.512916
token_balance{name="etherdelta",symbol="PPP",contract="0xc42209accc14029c1012fb5680d95fbd6036e2a0",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 1397693.222793370961529842
token_balance{name="etherdelta",symbol="PPT",contract="0xd4fa1460F537bb9085d22C7bcCB5DD450Ef28e3a",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 130051.41904102
token_balance{name="etherdelta",symbol="PRE",contract="0x88a3e4f35d64aad41a6d4030ac9afe4356cb84fa",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="PRG",contract="0x7728dfef5abd468669eb7f9b48a7f70a501ed29d",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 77659.522142
token_balance{name="etherdelta",symbol="PRL",contract="0x1844b21593262668b7248d0f57a220caaba46ab9",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 402840.335255952080372715
token_balance{name="etherdelta",symbol="PRO",contract="0x226bb599a12C826476e3A771454697EA52E9E220",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 113869.11356078
token_balance{name="etherdelta",symbol="PRON",contract="0xA3149E0fA0061A9007fAf307074cdCd290f0e2Fd",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="PRPS",contract="0x7641b2Ca9DDD58adDf6e3381c1F994Aac5f1A32f",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 18.902633090876177065
token_balance{name="etherdelta",symbol="PRS",contract="0x163733bcc28dbf26B41a8CfA83e369b5B3af741b",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 50163.617397539787055196
token_balance{name="etherdelta",symbol="PRSP",contract="0x0c04d4f331da8df75f9e2e271e3f3f1494c66c36",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="PT",contract="0x66497a283e0a007ba3974e837784c6ae323447de",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 342393.715926659253853997
token_balance{name="etherdelta",symbol="PTC",contract="0x2a8E98e256f32259b5E5Cb55Dd63C8e891950666",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="PTOY",contract="0x8Ae4BF2C33a8e667de34B54938B0ccD03Eb8CC06",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 49799.37157391
token_balance{name="etherdelta",symbol="PTWO",contract="0x5512e1d6a7be424b4323126b4f9e86d023f95764",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="PUC",contract="0xef6b4ce8c9bc83744fbcde2657b32ec18790458a",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="PXT",contract="0xc14830e53aa344e8c14603a91229a0b925b0b262",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 787359.42001708
token_balance{name="etherdelta",symbol="QAU",contract="0x671AbBe5CE652491985342e85428EB1b07bC6c64",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 16204.54281851
token_balance{name="etherdelta",symbol="QRL",contract="0x697beac28B09E122C4332D163985e8a73121b97F",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="QSP",contract="0x99ea4dB9EE77ACD40B119BD1dC4E33e1C070b80d",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 651239.487731342844179641
token_balance{name="etherdelta",symbol="QTQ",contract="0x2C3C1F05187dBa7A5f2Dd47Dca57281C4d4F183F",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="QTUM",contract="0x9a642d6b3368ddc662CA244bAdf32cDA716005BC",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0.128825959142667731
token_balance{name="etherdelta",symbol="RBLX",contract="0xfc2c4d8f95002c14ed0a7aa65102cac9e5953b5e",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 12265.077911872170150823
token_balance{name="etherdelta",symbol="RCN",contract="0xf970b8e36e23f7fc3fd752eea86f8be8d83375a6",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 752572.702733361445790371
token_balance{name="etherdelta",symbol="RDN",contract="0x255aa6df07540cb5d3d297f0d0d4d84cb52bc8e6",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 105418.703503806779873721
token_balance{name="etherdelta",symbol="REA",contract="0x767bA2915EC344015a7938E3eEDfeC2785195D05",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 973.92898845996754992
token_balance{name="etherdelta",symbol="REBL",contract="0x5f53f7a8075614b699baad0bc2c899f4bad8fbbf",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 394606.089993301575484666
token_balance{name="etherdelta",symbol="REN",contract="0x408e41876cCCDC0F92210600ef50372656052a38",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 243947.552121297301736148
token_balance{name="etherdelta",symbol="REP",contract="0xE94327D07Fc17907b4DB788E5aDf2ed424adDff6",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 1286.232715022289023009
token_balance{name="etherdelta",symbol="REQ",contract="0x8f8221aFbB33998d8584A2B05749bA73c37a938a",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 2323744.208066670433674025
token_balance{name="etherdelta",symbol="REX",contract="0xf05a9382A4C3F29E2784502754293D88b835109C",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 263305.342305804526439133
token_balance{name="etherdelta",symbol="RFR",contract="0xd0929d411954c47438dc1d871dd6081f5c5e149c",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 306808.3954
token_balance{name="etherdelta",symbol="RIPT",contract="0xdd007278b667f6bef52fd0a4c23604aa1f96039a",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 46705.59153881
token_balance{name="etherdelta",symbol="RLC",contract="0x607F4C5BB672230e8672085532f7e901544a7375",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 8647.656040356
token_balance{name="etherdelta",symbol="RLT",contract="0xcCeD5B8288086BE8c38E23567e684C3740be4D48",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 500.352
token_balance{name="etherdelta",symbol="RLTY",contract="0xbe99B09709fc753b09BCf557A992F6605D5997B0",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="RLX",contract="0x4a42d2c580f83dce404acad18dab26db11a1750e",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 67223516.824109499011680592
token_balance{name="etherdelta",symbol="RNDR",contract="0x0996bfb5d057faa237640e2506be7b4f9c46de0b",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 3466.677438463476411684
token_balance{name="etherdelta",symbol="ROCK",contract="0xA40106134c5bF4c41411554e6db99B95A15ed9d8",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 178.94240994992
token_balance{name="etherdelta",symbol="ROK",contract="0xc9de4b7f0c3d991e967158e4d4bfa4b51ec0b114",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="ROUND",contract="0x4993CB95c7443bdC06155c5f5688Be9D8f6999a5",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 2988.61461232
token_balance{name="etherdelta",symbol="RPL",contract="0xb4efd85c19999d84251304bda99e90b92300bd93",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 494618.024555305535205436
token_balance{name="etherdelta",symbol="RTN",contract="0x54b293226000ccBFC04DF902eEC567CB4C35a903",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="RVT",contract="0x3d1ba9be9f66b8ee101911bc36d3fb562eac2244",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 245354.943394575379878575
token_balance{name="etherdelta",symbol="S-A-PAT",contract="0x1ec8fe51a9b6a3a6c427d17d9ecc3060fbc4a45c",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="S-ETH",contract="0x3eb91d237e491e0dee8582c402d85cb440fb6b54",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="SALT",contract="0x4156D3342D5c385a87D264F90653733592000581",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 146470.49633509
token_balance{name="etherdelta",symbol="SAN",contract="0x7C5A0CE9267ED19B22F8cae653F198e3E8daf098",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 25893.063385165872951209
token_balance{name="etherdelta",symbol="SCANDI",contract="0x78fe18e41f436e1981a3a60d1557c8a7a9370461",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="SCL",contract="0xd7631787b4dcc87b1254cfd1e5ce48e96823dee8",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 140080.80447185
token_balance{name="etherdelta",symbol="SENSE",contract="0x6745fAB6801e376cD24F03572B9C9B0D4EdDDCcf",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 148415.35197445
token_balance{name="etherdelta",symbol="SET",contract="0xe06eda7435ba749b047380ced49121dde93334ae",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="SEXY",contract="0x98f5e9b7f0e33956c0443e81bf7deb8b5b1ed545",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 48000112.0
token_balance{name="etherdelta",symbol="SGEL",contract="0xa1ccc166faf0E998b3E33225A1A0301B1C86119D",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="SGT",contract="0xd248B0D48E44aaF9c49aea0312be7E13a6dc1468",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="SHIT",contract="0xEF2E9966eb61BB494E5375d5Df8d67B7dB8A780D",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 219072
token_balance{name="etherdelta",symbol="SIFT",contract="0x8a187d5285d316bcbc9adafc08b51d70a0d8e000",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 871
token_balance{name="etherdelta",symbol="SIG",contract="0x6888a16eA9792c15A4DCF2f6C623D055c8eDe792",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 2094938.451475935192039397
token_balance{name="etherdelta",symbol="SKIN",contract="0x2bDC0D42996017fCe214b21607a515DA41A9E0C5",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 1352527.463613
token_balance{name="etherdelta",symbol="SKO1",contract="0x4994e81897a920c0FEA235eb8CEdEEd3c6fFF697",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="SKR",contract="0x4c382F8E09615AC86E08CE58266CC227e7d4D913",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 249580.675423
token_balance{name="etherdelta",symbol="SKRP (Phase 1-E)",contract="0x324a48ebcbb46e61993931ef9d35f6697cd2901b",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="SKRP (Phase 1)",contract="0x6E34d8d84764D40f6D7b39cd569Fd017bF53177D",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="SKRP",contract="0xfdFE8b7aB6CF1bD1E3d14538Ef40686296C42052",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 138655.357034612681508387
token_balance{name="etherdelta",symbol="SLT",contract="0x7A5fF295Dc8239d5C2374E4D894202aAF029Cab6",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="SMART",contract="0x6F6DEb5db0C4994A8283A01D6CFeEB27Fc3bBe9C",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 112525
token_balance{name="etherdelta",symbol="SMT (SmartMesh)",contract="0x55f93985431fc9304077687a35a1ba103dc1e081",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 1.933
token_balance{name="etherdelta",symbol="SMT (Smart Node)",contract="0x2dcfaac11c9eebd8c6c42103fe9e2a6ad237af27",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 25.0
token_balance{name="etherdelta",symbol="SNC",contract="0xF4134146AF2d511Dd5EA8cDB1C4AC88C57D60404",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 176375.581862621039184495
token_balance{name="etherdelta",symbol="SND",contract="0xf333b2Ace992ac2bBD8798bF57Bc65a06184afBa",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 115602
token_balance{name="etherdelta",symbol="SNG",contract="0xcFD6Ae8BF13f42DE14867351eAff7A8A3b9FbBe7",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 124647.81704197
token_balance{name="etherdelta",symbol="SNGLS",contract="0xaeC2E87E0A235266D9C5ADc9DEb4b2E29b54D009",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 290958
token_balance{name="etherdelta",symbol="SNIP",contract="0x44F588aEeB8C44471439D1270B3603c66a9262F1",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 169147440.696030500110220051
token_balance{name="etherdelta",symbol="SNM",contract="0x983F6d60db79ea8cA4eB9968C6aFf8cfA04B3c63",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 435536.856641117676760652
token_balance{name="etherdelta",symbol="SNOV",contract="0xbdc5bac39dbe132b1e030e898ae3830017d7d969",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 5292494.893140348222479154
token_balance{name="etherdelta",symbol="SNT",contract="0x744d70FDBE2Ba4CF95131626614a1763DF805B9E",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 1362162.461533406056762009
token_balance{name="etherdelta",symbol="SOL",contract="0x1f54638b7737193ffd86c19ec51907a7c41755d8",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 420947.287166
token_balance{name="etherdelta",symbol="SPANK",contract="0x42d6622dece394b54999fbd73d108123806f6a18",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 3332064.455235052182421323
token_balance{name="etherdelta",symbol="SPARC",contract="0x58bf7df57d9DA7113c4cCb49d8463D4908C735cb",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 2139.395263314917
token_balance{name="etherdelta",symbol="SPARTA",contract="0x24aef3bf1a47561500f9430d74ed4097c47f51f2",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 1360.174
token_balance{name="etherdelta",symbol="SPF",contract="0x85089389C14Bd9c77FC2b8F0c3d1dC3363Bf06Ef",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 6129.766980945320731577
token_balance{name="etherdelta",symbol="SPN",contract="0x20F7A3DdF244dc9299975b4Da1C39F8D5D75f05A",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="SRN",contract="0x68d57c9a1C35f63E2c83eE8e49A64e9d70528D25",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 22152.480149651947138861
token_balance{name="etherdelta",symbol="SS",contract="0xB15fE5a123e647ba594CEa7A1E648646f95EB4AA",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 10.0
token_balance{name="etherdelta",symbol="STAC",contract="0x9a005c9a89bd72a4bd27721e7a09a3c11d2b03c4",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 11941.552168791638768062
token_balance{name="etherdelta",symbol="STAR",contract="0xF70a642bD387F94380fFb90451C2c81d4Eb82CBc",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 5509422.625171152964158815
token_balance{name="etherdelta",symbol="STC",contract="0x629aEe55ed49581C33ab27f9403F7992A289ffd5",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="STK",contract="0xaE73B38d1c9A8b274127ec30160a4927C4d71824",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 366004.159161837744505628
token_balance{name="etherdelta",symbol="STN",contract="0x599346779e90fc3F5F997b5ea715349820F91571",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="STORJ",contract="0xB64ef51C888972c908CFacf59B47C1AfBC0Ab8aC",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 39084.29980089
token_balance{name="etherdelta",symbol="STORM",contract="0xD0a4b8946Cb52f0661273bfbC6fD0E0C75Fc6433",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 2006399.751925904354069354
token_balance{name="etherdelta",symbol="STP",contract="0xecd570bBf74761b960Fa04Cc10fe2c4e86FfDA36",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 251013.70775151
token_balance{name="etherdelta",symbol="STRC",contract="0x46492473755e8dF960F8034877F61732D718CE96",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 2434.36003696
token_balance{name="etherdelta",symbol="STX",contract="0x006BeA43Baa3f7A6f765F14f10A1a1b08334EF45",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 91985.075887898712836853
token_balance{name="etherdelta",symbol="SUB",contract="0x12480E24eb5bec1a9D4369CaB6a80caD3c0A377A",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 882629.04
token_balance{name="etherdelta",symbol="SWM",contract="0x9e88613418cf03dca54d6a2cf6ad934a78c7a17a",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 58328.669899539002485199
token_balance{name="etherdelta",symbol="SWT",contract="0xB9e7F8568e08d5659f5D29C4997173d84CdF2607",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 7775.548683164967220735
token_balance{name="etherdelta",symbol="SXDT",contract="0x12b306fa98f4cbb8d4457fdff3a0a0a56f07ccdf",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 2462459.56676967184587042
token_balance{name="etherdelta",symbol="SXUT",contract="0x2c82c73d5b34aa015989462b2948cd616a37641f",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 294413.326923029597723296
token_balance{name="etherdelta",symbol="SYN",contract="0x10b123fddde003243199aad03522065dc05827a0",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="SenSatorI",contract="0x4ca74185532dc1789527194e5b9c866dd33f4e82",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 5200.0
token_balance{name="etherdelta",symbol="TAU",contract="0xc27a2f05fa577a83ba0fdb4c38443c0718356501",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 2831642.95899992944838963
token_balance{name="etherdelta",symbol="TBC2",contract="0xFACCD5Fc83c3E4C3c1AC1EF35D15adf06bCF209C",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="TBT",contract="0xAFe60511341a37488de25Bef351952562E31fCc1",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 113055.68415179
token_balance{name="etherdelta",symbol="TEL",contract="0x85e076361cc813a908ff672f9bad1541474402b2",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 53923390.66
token_balance{name="etherdelta",symbol="TFL",contract="0xa7f976C360ebBeD4465c2855684D1AAE5271eFa9",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 33705.29789313
token_balance{name="etherdelta",symbol="THETA",contract="0x3883f5e181fccaf8410fa61e12b59bad963fb645",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 177762.109406112638059649
token_balance{name="etherdelta",symbol="TIME",contract="0x6531f133e6DeeBe7F2dcE5A0441aA7ef330B4e53",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 169.40507972
token_balance{name="etherdelta",symbol="TIO",contract="0x80bc5512561c7f85a3a9508c7df7901b370fa1df",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 21984.403477554118318042
token_balance{name="etherdelta",symbol="TIX",contract="0xEa1f346faF023F974Eb5adaf088BbCdf02d761F4",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 169497.147780291201121716
token_balance{name="etherdelta",symbol="TKN",contract="0xaAAf91D9b90dF800Df4F55c205fd6989c977E73a",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 38080.45957405
token_balance{name="etherdelta",symbol="TNT",contract="0x08f5a9235b08173b7569f83645d2c7fb55e8ccd8",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 1670376.55837181
token_balance{name="etherdelta",symbol="TRC",contract="0xcB3F902bf97626391bF8bA87264bbC3DC13469be",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="TRCN",contract="0x566Fd7999B1Fc3988022bD38507A48F0bCf22c77",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="TRST",contract="0xcb94be6f13a1182e4a4b6140cb7bf2025d28e41b",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 16171.379994
token_balance{name="etherdelta",symbol="TRX",contract="0xf230b790e05390fc8295f4d3f60332c93bed42e2",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 22240959.014818
token_balance{name="etherdelta",symbol="TWN",contract="0x2eF1aB8a26187C58BB8aAeB11B2fC6D25C5c0716",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 701733.52855
token_balance{name="etherdelta",symbol="TWNKL",contract="0xfbd0d1c77b501796a35d86cf91d65d9778eee695",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 831234.101
token_balance{name="etherdelta",symbol="TaaS",contract="0xE7775A6e9Bcf904eb39DA2b68c5efb4F9360e08C",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 1828.609356
token_balance{name="etherdelta",symbol="UKG",contract="0x24692791bc444c5cd0b81e3cbcaba4b04acd1f3b",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 160408.425506813478636426
token_balance{name="etherdelta",symbol="UQC",contract="0xd01db73e047855efb414e6202098c4be4cd2423b",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 523.417
token_balance{name="etherdelta",symbol="USDT",contract="0xdac17f958d2ee523a2206206994597c13d831ec7",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="UTK",contract="0x70a72833d6bf7f508c8224ce59ea1ef3d0ea3a38",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 609184.396531256367396285
token_balance{name="etherdelta",symbol="UTN-P",contract="0x9e3319636e2126e3c0bc9e3134AEC5e1508A46c7",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="Unicorn",contract="0x89205A3A3b2A69De6Dbf7f01ED13B2108B2c43e7",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="VDOC",contract="0x82BD526bDB718C6d4DD2291Ed013A5186cAE2DCa",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="VEE",contract="0x340d2bde5eb28c1eed91b2f790723e3b160613b7",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 9078514.580391640454398004
token_balance{name="etherdelta",symbol="VENUS",contract="0xEbeD4fF9fe34413db8fC8294556BBD1528a4DAca",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="VERI",contract="0x8f3470A7388c05eE4e7AF3d01D8C722b0FF52374",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 22200.940679658225779464
token_balance{name="etherdelta",symbol="VEN",contract="0xD850942eF8811f2A866692A623011bDE52a462C1",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 106979.805588890163923479
token_balance{name="etherdelta",symbol="VIB",contract="0x2C974B2d0BA1716E644c1FC59982a89DDD2fF724",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 149926.51132404012360395
token_balance{name="etherdelta",symbol="VIBE",contract="0xe8ff5c9c75deb346acac493c463c8950be03dfba",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 1073673.844994362690173027
token_balance{name="etherdelta",symbol="VIBEX",contract="0x882448f83d90b2bf477af2ea79327fdea1335d93",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="VIT",contract="0x23b75Bc7AaF28e2d6628C3f424B3882F8f072a3c",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 20636.2
token_balance{name="etherdelta",symbol="VIU",contract="0x519475b31653e46d20cd09f9fdcf3b12bdacb4f5",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 3732000.992453165007085347
token_balance{name="etherdelta",symbol="VOC",contract="0xc3bc9eb71f75ec439a6b6c8e8b746fcf5b62f703",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="VOISE",contract="0x83eEA00D838f92dEC4D1475697B9f4D3537b56E3",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 1601274.39460573
token_balance{name="etherdelta",symbol="VRS",contract="0xeDBaF3c5100302dCddA53269322f3730b1F0416d",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="VSL",contract="0x5c543e7AE0A1104f78406C340E9C64FD9fCE5170",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 20422.209226633031398405
token_balance{name="etherdelta",symbol="WAX",contract="0x39Bb259F66E1C59d5ABEF88375979b4D20D98022",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 405946.80242411
token_balance{name="etherdelta",symbol="WBA",contract="0x74951B677de32D596EE851A233336926e6A2cd09",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 199294.0884575
token_balance{name="etherdelta",symbol="WCT",contract="0x6a0a97e47d15aad1d132a1ac79a480e3f2079063",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="WETH",contract="0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="WHEN",contract="0xf4fe95603881d0e07954fd7605e0e9a916e42c44",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 2875.0
token_balance{name="etherdelta",symbol="WHO",contract="0xe933c0Cd9784414d5F278C114904F5A84b396919",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="WIC",contract="0x62cd07d414ec50b68c7ecaa863a23d344f2d062f",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="WILD",contract="0xD3C00772B24D997A812249ca637a921e81357701",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 638514.235337219693405353
token_balance{name="etherdelta",symbol="WINGS",contract="0x667088b212ce3d06a1b553a7221E1fD19000d9aF",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 4370.339026339096956122
token_balance{name="etherdelta",symbol="WLK",contract="0xF6B55acBBC49f4524Aa48D19281A9A77c54DE10f",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 21266.986134957621659181
token_balance{name="etherdelta",symbol="WOLK",contract="0x728781E75735dc0962Df3a51d7Ef47E798A7107E",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 258165.847870047816780925
token_balance{name="etherdelta",symbol="WPC",contract="0x62087245087125d3db5b9a3d713d78e7bbc31e54",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="WPR",contract="0x4CF488387F035FF08c371515562CBa712f9015d4",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 514577.727323924156909794
token_balance{name="etherdelta",symbol="WRK",contract="0x71e8d74ff1c923e369d0e70dfb09866629c4dd35",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="WYV",contract="0x056017c55aE7AE32d12AeF7C679dF83A85ca75Ff",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 12863.415543436965199175
token_balance{name="etherdelta",symbol="WaBi",contract="0x286BDA1413a2Df81731D4930ce2F862a35A609fE",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 111120.242280957550843165
token_balance{name="etherdelta",symbol="WiC",contract="0x5e4ABE6419650CA839Ce5BB7Db422b881a6064bB",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 1223593.314484240030042618
token_balance{name="etherdelta",symbol="X8X",contract="0x910Dfc18D6EA3D6a7124A6F8B5458F281060fa4c",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 469823.885500842555371336
token_balance{name="etherdelta",symbol="XAUR",contract="0x4DF812F6064def1e5e029f1ca858777CC98D2D81",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 69595.81620641
token_balance{name="etherdelta",symbol="XBP",contract="0x28dee01d53fed0edf5f6e310bf8ef9311513ae40",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 458417.287298225212470351
token_balance{name="etherdelta",symbol="XCC",contract="0x4d829f8c92a6691c56300d020c9e0db984cfe2ba",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 1.0
token_balance{name="etherdelta",symbol="XGM",contract="0x533ef0984b2FAA227AcC620C67cce12aA39CD8CD",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 2197
token_balance{name="etherdelta",symbol="XGT",contract="0x30f4A3e0aB7a76733D8b60b89DD93c3D0b4c9E2f",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="XID",contract="0xB110eC7B1dcb8FAB8dEDbf28f53Bc63eA5BEdd84",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="XNK",contract="0xBC86727E770de68B1060C91f6BB6945c73e10388",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 132125.136462750227985186
token_balance{name="etherdelta",symbol="XNN",contract="0xab95e915c123fded5bdfb6325e35ef5515f1ea69",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 1428308.475618502556030014
token_balance{name="etherdelta",symbol="XNT",contract="0x572e6f318056ba0c5d47a422653113843d250691",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="XRL",contract="0xB24754bE79281553dc1adC160ddF5Cd9b74361a4",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 3049755.594327605
token_balance{name="etherdelta",symbol="XSC",contract="0x0F513fFb4926ff82D7F60A05069047AcA295C413",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 675870.022020436646396172
token_balance{name="etherdelta",symbol="YUPIE",contract="0x0F33bb20a282A7649C7B3AFf644F084a9348e933",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 17980.155559940784326754
token_balance{name="etherdelta",symbol="ZAP",contract="0x6781a0f84c7e9e846dcb84a9a5bd49333067b104",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 823148.447415897000724906
token_balance{name="etherdelta",symbol="ZIL",contract="0x05f4a42e251f2d52b8ed15E9FEdAacFcEF1FAD27",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 7537136.442411115073
token_balance{name="etherdelta",symbol="ZRX",contract="0xE41d2489571d322189246DaFA5ebDe1F4699F498",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 1435303.751046955446524494
token_balance{name="etherdelta",symbol="ZST",contract="0xe386b139ed3715ca4b18fd52671bdcea1cdfe4b1",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_balance{name="etherdelta",symbol="cV",contract="0xdA6cb58A0D0C01610a29c5A65c303e13e885887C",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 17464924.37305065447721005
token_balance{name="etherdelta",symbol="eBCH",contract="0xafc39788c51f0c1ff7b55317f3e70299e521fff6",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 1048667.68989734
token_balance{name="etherdelta",symbol="eBTC",contract="0xeb7c20027172e5d143fb030d50f91cece2d1485d",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 2099877.50257294
token_balance{name="etherdelta",symbol="eGAS",contract="0xb53a96bcbdd9cf78dff20bab6c2be7baec8f00f8",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 2658390.96282287
token_balance{name="etherdelta",symbol="onG",contract="0xd341d1680eeee3255b8c4c75bcce7eb57f144dae",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 509226.564895201084026963
token_balance{name="etherdelta",symbol="YEED",contract="0x6f7a4bac3315b5082f793161a22e26666d22717f",address="0x8d12A197cB00D4747a1fe03395095ce2A5CC6819"} 0
token_query_seconds{name="all"} 10.037673145
```

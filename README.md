# 啄木鸟-基于cosmos的电子病历管理系统

## 开始

```shell
git clone https://github.com/xwjahahahaha/woodpecker.git
starport serve -v
```

<!-- more -->

实现链上的电子医疗病历的CRUD， 解决电子病历链上共享与隐私保护。

## 展示

**（仓库仅cosmos模块代码，但提供了接口可自行实现前端）**

![tNlRRf](http://xwjpics.gumptlu.work/qinniu_uPic/tNlRRf.png)

![xyG4Ay](http://xwjpics.gumptlu.work/qinniu_uPic/xyG4Ay.png)

链上登陆

![H0ifiT](http://xwjpics.gumptlu.work/qinniu_uPic/H0ifiT.png)

![fKMOSG](http://xwjpics.gumptlu.work/qinniu_uPic/fKMOSG.png)

![6wvFjD](http://xwjpics.gumptlu.work/qinniu_uPic/6wvFjD.png)

![MatrUX](http://xwjpics.gumptlu.work/qinniu_uPic/MatrUX.png)

![MXiIV8](http://xwjpics.gumptlu.work/qinniu_uPic/MXiIV8.png)

![IhZQmm](http://xwjpics.gumptlu.work/qinniu_uPic/IhZQmm.png)

## 使用

病历被分为三部分存储Attribute、MedicalHistory、BodyIndex

hashkey字段为病历唯一key，Attribute、BodyIndex没有创建create直接set即可，所以没有ID（没有创建就是创建，创建过了就是修改）

MedicalHistory创建需要使用create，其有ID字段，即表示一个病人的病历会有多个

### Attribute

#### 1.cli客户端

##### 创建:

```shell
$ woodpeckercli tx woodpecker set-attribute [name] [idNumber] [address] [job] [nation] [hashKey] [flags] --from user
```

```shell
$ woodpeckercli tx woodpecker set-attribute xwj 123 cq it china 1902 --from user1 -y
```

![Ag8dab](http://xwjpics.gumptlu.work/qinniu_uPic/Ag8dab.png)

##### 查询：

all：

```shell
$ woodpeckercli query woodpecker list-attribute
```

![pwUJ4H](http://xwjpics.gumptlu.work/qinniu_uPic/pwUJ4H.png)

one：

```shell
$ woodpeckercli query woodpecker list-attribute [hashKey]
```

![vH4D7Q](http://xwjpics.gumptlu.work/qinniu_uPic/vH4D7Q.png)

##### 修改

```shell
$ woodpeckercli tx woodpecker set-attribute [name] [idNumber] [address] [job] [nation] [hashKey] [flags] --from user
```

同创建

##### 删除

```shell
$ woodpeckercli tx woodpecker delete-attribute [hashKey]
```

#### 2.Rest接口

```shell
# 查
# one
your IP:1317/woodpecker/attribute/[hashKey]
# all
your IP:1317/woodpecker/attribute
```

交易通过接口无法签名，可以通过前端js实现[cosmosJs](https://github.com/cosmostation/cosmosjs)

### MedicalHistory

#### 1.cli客户端

##### 创建

![fGXayh](http://xwjpics.gumptlu.work/qinniu_uPic/fGXayh.png)

```shell
$ woodpeckercli tx woodpecker create-medicalHistory 协和医院 2021/03/14 精神科 脑子 小阳 深井冰 放弃治疗 1902 --from user1 -y
```

![RWATVn](http://xwjpics.gumptlu.work/qinniu_uPic/RWATVn.png)

##### 查询

###### 查询所有人的所有病历

![xxzvdR](http://xwjpics.gumptlu.work/qinniu_uPic/xxzvdR.png)

```shell
$ woodpeckercli query woodpecker list-all-medicalHistory
```

![kDTyct](http://xwjpics.gumptlu.work/qinniu_uPic/kDTyct.png)

###### 查询一个人的所有病历

![OS8Rzj](http://xwjpics.gumptlu.work/qinniu_uPic/OS8Rzj.png)

```shell
$ woodpeckercli query woodpecker list-medicalHistory 1902
```

1902:

![xTsVU0](http://xwjpics.gumptlu.work/qinniu_uPic/xTsVU0.png)

xwj:

![AM0Byh](http://xwjpics.gumptlu.work/qinniu_uPic/AM0Byh.png)

###### 查询一个人的单个病历

```shell
$ woodpeckercli query woodpecker get-medicalHistory [hashKey] [id]
```

```shell
$ woodpeckercli query woodpecker get-medicalHistory 1902 0
```

![N7I2Ri](http://xwjpics.gumptlu.work/qinniu_uPic/N7I2Ri.png)

##### 修改

###### 修改一个人的某个病历

```shell
$ woodpeckercli tx woodpecker set-medicalHistory [id]  [medicalInstitutionName] [diagnosisTime] [diagnosisDepartment] [diagnosisType] [diagnosisDoctor] [diagnosisResult] [treatmentOptions] [hashKey] [flags]
```

```shell
$ woodpeckercli tx woodpecker set-medicalHistory 0 北京医院 2021/4/10 骨科 骨质疏松 小阳 骨头坏死 截肢 1902 --from user1 -y
```

再次查询:

![6uYHfr](http://xwjpics.gumptlu.work/qinniu_uPic/6uYHfr.png)

##### 删除

```shell
$ woodpeckercli tx woodpecker delete-medicalHistory [id] [hashKey]
```

```shell
$ woodpeckercli tx woodpecker delete-medicalHistory 4 xwj --from user1 -y
```

再次查询;

![KlBMU4](http://xwjpics.gumptlu.work/qinniu_uPic/KlBMU4.png)

#### 2.rest接口测试

注意：需要交易的方法（创建、修改、删除）都需要本地签名才会生效, 建议前端使用cosmosjs

```shell
# Then sign this transaction
# NOTE: In a real environment the raw transaction should be signed on the client side. Also the sequence needs to be adjusted, depending on what the query of user2's account has shown.
woodpeckercli tx sign unsignedTx.json --from user1 --offline --chain-id namechain --sequence 1 --account-number 0 > signedTx.json

# And finally broadcast the signed transaction
woodpeckercli tx broadcast signedTx.json
```

##### 创建

```yaml
URL : IP Address:1317/woodpecker/medicalHistory
POST
Data Simple:

{
	"base_req": {
		"from": "cosmos15qfsrthwsfu378m5epe2fzggsu5m7r7d0yexep",
		"chain_id": "woodpacker"
	},
    "creator": "cosmos15qfsrthwsfu378m5epe2fzggsu5m7r7d0yexep",
    "medicalInstitutionName": "南京和谐",
    "diagnosisTime": "2021/08/24",
    "diagnosisDepartment": "dkaskd",
    "diagnosisType": "脑壳",
    "diagnosisDoctor": "小样",
    "diagnosisResult": "deal",
    "treatmentOptions": "ldsl",
    "hash_key": "xwj"
}
```

##### 查询

###### 查询所有人的所有病历

```yaml
URL: IP Address:1317/woodpecker/medicalHistory
GET
```

result simple:

```json
{
    "height": "0",
    "result": [
        {
            "creator": "cosmos1r304rwx7z4vh6yz4wkp4utwchg03ymc6rw45je",
            "id": "0",
            "medicalInstitutionName": "协和医院",
            "diagnosisTime": "2021/03/14",
            "diagnosisDepartment": "精神科",
            "diagnosisType": "脑子",
            "diagnosisDoctor": "小阳",
            "diagnosisResult": "深井冰",
            "treatmentOptions": "放弃治疗"
        },
        {
            "creator": "cosmos1r304rwx7z4vh6yz4wkp4utwchg03ymc6rw45je",
            "id": "1",
            "medicalInstitutionName": "协和医院",
            "diagnosisTime": "2021/03/14",
            "diagnosisDepartment": "精神科",
            "diagnosisType": "脑子",
            "diagnosisDoctor": "小阳",
            "diagnosisResult": "深井冰",
            "treatmentOptions": "放弃治疗"
        }
    ]
}
```

###### 查询一个人的所有病历

```yaml
URL: IP Address:1317/woodpecker/medicalHistory/[hashKey]
GET
```

###### 查询一个人的单个病历

```yaml
URL: IP Address:1317/woodpecker/medicalHistory/[hashKey]/[id]
GET
```


##### 修改

###### 修改一个人的某个病历

```yaml
URL: IP Address:1317/woodpecker/medicalHistory
PUT

Data Sample:
{
	"base_req": {
		"from": "cosmos1x7acsyd53mmjc3znwt8xw2ju3vpwm6gku7muu0",
		"chain_id": "woodpacker"
	},
    "creator": "cosmos1x7acsyd53mmjc3znwt8xw2ju3vpwm6gku7muu0",
    "medicalInstitutionName": "xxxx",
    "diagnosisTime": "xxxxxs",
    "diagnosisDepartment": "dkaskd",
    "diagnosisType": "xxxxx",
    "diagnosisDoctor": "小样",
    "diagnosisResult": "deal",
    "treatmentOptions": "ldsl",
    "hash_key": "xwj",
    "id": "3"
}
```

##### 删除

```yaml
URL: IP Address:1317/woodpecker/medicalHistory
DELET

Data Sample:
{
	"base_req": {
		"from": "cosmos1x7acsyd53mmjc3znwt8xw2ju3vpwm6gku7muu0",
		"chain_id": "woodpacker"
	},
    "creator": "cosmos1x7acsyd53mmjc3znwt8xw2ju3vpwm6gku7muu0",
    "hash_key": "xwj",
    "id": "3"
}
```

### BodyIndex

参数较长建议使用文件

#### 1.cli

##### 创建/更新

```shell
$ woodpeckercli tx woodpecker set-bodyIndex [age] [sex] [nation] [weight] [height] [weightIndex] [obesityWaistline] [waistline] [maxBloodPressure] [minBloodPressure] [goodCholesterol] [batCholesterol] [totalCholesterol] [Dyslipidemia] [pvd] [sportActivities] [education] [marry] [income] [sourceCase] [visionBad] [drink] [highBloodPressure] [familialHighBloodPressure] [diabetes] [familialDiabetes] [hepatitis] [familialHepatitis] [chronicFatigue] [alf] [hashKey]
```

```shell
$ woodpeckercli tx woodepecker set-bodyIndex $(<data.txt)
```

##### 查询

```shell
# all
$ woodpeckercli query woodpecker list-bodyIndex
# one
$ woodpeckercli query woodpecker list-bodyIndex [hashKey]
```

##### 删除

```shell
$ woodpeckercli tx woodpecker delete-bodyIndex [hashKey]
```

#### 2.Rest

##### 查询

```shell
# all
IP:1317/woodpecker/attribute
# one
IP:1317/woodpecker/attribute/[hashKey]
```


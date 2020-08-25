# excel2json
Copyright (C) 2020 by donggeuri0320 donggeuri0320@gmail.com

Excel2json is a converting script that supports to manage well structured excel data to json format.

You can freely redistribute this product with A-CUP-OF-BEER License (See source code)

It's similar to https://github.com/coolengineer/excel2json

# USAGE

#### Excel2json converts all sheets of all .xlsx files in the xlsx folder. The converted contents are created in the ./json folder.

##### A type must be specified above the name.
<pre>
n: number
narr: number array
s: string
sarr: string array

Please refer to ./xlsx/test.xlsx
</pre>

## Simple Object
||A|B|C
---- | ---- | ---- | ----
1 | | s | s 
2| initdata | $key | $value 
3| | a | b 
<pre>
{
	"initdata":{
		"a": "apple",
		"b": "banana"
	}
}
</pre>
## Objects in Object
|| A | B | C | D
---- | ---- | ---- | ---- | ----
1||s|n|narr
2|girlgroup |	$key	| peopleCnt	| peopleIds
3||blackpink	|5|	1~3, 12
4||twice	|9|	5~11,13
5||hmm..	|||||	

An integer can be specified as a range value.

<pre>
{
  "girlgroup":{
		"blackpink":{"peopleCnt":5, "peopleIds":[1,2,3,12]},
		"twice":{"peopleCnt":9, "peopleIds":[5,6,7,8,9,10,11,13]},
		"hmm..":{"peopleCnt":0, "peopleIds":[]}
	}
}
</pre>

## Object Array

||A|B|C|D
---- | ---- | ---- | ---- | ----
1||n|s|sarr
2|people |	peopleId	|name|	test
3||	1|	jisso|	a,b,c
4||	2|	jennie|	a
5||	3|	rose|	|
6||	4|	나연|	b
7||	5|	정연|	c
8||	6|	모모|	d, e
9||	7|	사나|	|
10||	8|	지효|	|
11||	9|	미나|	|
12||	10|	다현|	f
13||	11|	채영|	|
14||	12|	lisa|	|
15||	13|	쯔위| |

<pre>
{
  "people":[
      {"peopleId":1, "name":"jisso", "test":["a","b","c"]},
      {"peopleId":2, "name":"jennie", "test":["a"]},
      {"peopleId":3, "name":"rose", "test":[]},
      {"peopleId":4, "name":"나연", "test":["b"]},
      {"peopleId":5, "name":"정연", "test":["c"]},
      {"peopleId":6, "name":"모모", "test":["d","e"]},
      {"peopleId":7, "name":"사나", "test":[]},
      {"peopleId":8, "name":"지효", "test":[]},
      {"peopleId":9, "name":"미나", "test":[]},
      {"peopleId":10, "name":"다현", "test":["f"]},
      {"peopleId":11, "name":"채영", "test":[]},
      {"peopleId":12, "name":"lisa", "test":[]},
      {"peopleId":13, "name":"쯔위", "test":[]}
    ]
}
</pre>

##### Run excel2json.bat

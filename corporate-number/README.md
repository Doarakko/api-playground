# corporate-number

## Example

```sh
docker exec -it corporate-number-app python main.py 中日ドラゴンズ
```

```json
{
    "corporations": {
        "lastUpdateDate": "2023-02-17",
        "count": "1",
        "divideNumber": "1",
        "divideSize": "1",
        "corporation": {
            "sequenceNumber": "1",
            "corporateNumber": "1180001037972",
            "process": "12",
            "correct": "0",
            "updateDate": "2018-12-28",
            "changeDate": "2018-12-25",
            "name": "株式会社中日ドラゴンズ",
            "nameImageId": null,
            "kind": "301",
            "prefectureName": "愛知県",
            "cityName": "名古屋市東区",
            "streetNumber": "大幸南１丁目１番５１号",
            "addressImageId": null,
            "prefectureCode": "23",
            "cityCode": "102",
            "postCode": "4610047",
            "addressOutside": null,
            "addressOutsideImageId": null,
            "closeDate": null,
            "closeCause": null,
            "successorCorporateNumber": null,
            "changeCause": null,
            "assignmentDate": "2015-10-05",
            "latest": "1",
            "enName": null,
            "enPrefectureName": null,
            "enCityName": null,
            "enAddressOutside": null,
            "furigana": "チュウニチドラゴンズ",
            "hihyoji": "0"
        }
    }
}
```

## Requirements

- Docker Compose
- 法人番号システムWeb-API アプリケーションID

## Usage

```sh
usage: main.py [-h] [-m {1,2}] [-t {1,2,3}] [-c {1,2}] q

各引数、レスポンスの詳細な仕様はこちらを確認してください。https://www.houjin-bangou.nta.go.jp/documents/k-web-api-kinou-gaiyo.pdf

positional arguments:
  q           検索クエリ

optional arguments:
  -h, --help  show this help message and exit
  -m {1,2}    検索方式 1: 前方一致, 2: 部分一致
  -t {1,2,3}  検索対象 1: あいまい, 2: 完全一致, 3: 英語表記登録情報
  -c {1,2}    変更履歴 1: 含める, 2: 含めない
```

### 1. Setup

```sh
git clone https://github.com/Doarakko/api-playground
cd api-playground/corporate-number
cp .env.example .env
```

Enter your application id.

```.env
APP_ID=abcd
```

### 2. Build and start

```sh
docker-compose up --build -d
```

### 3. Run

```sh
docker exec -it corporate-number-app python main.py <q>
```

### 3. Check

`./app/output/<q>.json`

## Reference

- [法人番号システム Web-API](https://www.houjin-bangou.nta.go.jp/webapi/)

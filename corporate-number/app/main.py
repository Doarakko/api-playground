import argparse
import json
import os

import mojimoji
import requests
import xmltodict


def run(q: str, mode: int, target: int, change: int):
    name = mojimoji.han_to_zen(q)

    params = {
        "id": os.getenv("APP_ID"),
        "name": name,
        "type": "12",
        "mode": mode,
        "target": target,
        "change": change,
    }

    try:
        response = requests.get(
            "https://api.houjin-bangou.nta.go.jp/4/name", params=params
        )
        response.raise_for_status()
    except requests.exceptions.RequestException as e:
        print(e.response.text)
        raise e

    d = xmltodict.parse(response.text)

    path = f"output/{q}.json"
    with open(path, "w") as f:
        json.dump(d, f, indent=4, ensure_ascii=False)

    print(f"saved to {path}")


if __name__ == "__main__":
    parser = argparse.ArgumentParser(
        description="各引数、レスポンスの詳細な仕様はこちらを参照してください。https://www.houjin-bangou.nta.go.jp/documents/k-web-api-kinou-gaiyo.pdf"
    )

    parser.add_argument("q", type=str, help="検索クエリ")
    parser.add_argument(
        "-m", type=int, choices=[1, 2], default=1, help="検索方式 1: 前方一致, 2: 部分一致"
    )
    parser.add_argument(
        "-t",
        type=int,
        choices=[1, 2, 3],
        default=1,
        help="検索対象 1: あいまい, 2: 完全一致, 3: 英語表記登録情報",
    )
    parser.add_argument(
        "-c", type=int, choices=[1, 2], default=0, help="変更履歴 1: 含める, 2: 含めない"
    )

    args = parser.parse_args()

    run(args.q, args.m, args.t, args.c)

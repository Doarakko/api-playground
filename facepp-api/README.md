# Face++ Detect API
![image.png](https://qiita-image-store.s3.amazonaws.com/0/245792/b8de8352-0e4d-6e99-e606-325a105ac94e.png)

[Face++ Detect API](https://www.faceplusplus.com/face-detection/)を使用して, ローカル画像から顔を検出し, 顔の切り取りを行います.  
顔が正面を向いていて, 一定サイズ以上大きい顔画像のみを取得することを目的とにします.  
そのため, 以下の条件を満たす顔画像は顔の切り取りは行いません.

- 検出した顔の数が0または複数
- 顔の横揺れの角度が15°以上または-15°以下
- 顔の縦揺れの角度が10°以上または-10°以下
- 瞳孔間の距離が40ピクセル未満

取得した顔の特徴量を後で使用することを想定するため, レスポンスをJSON形式で保存しました.  
その後, JSONファイルをロードして顔の切り取りを行います.

## Requirement
- Python 3
    - opencv-python 3.4.0.12
- Face++ Detect API v3.0


## Usage
1. Clone code
```
$ git clone https://github.com/Doarakko/api-challenge
$ cd facepp-api-clip-face
```

2. Enter your API KEY and API SECRET
![facepp.jpg](https://qiita-image-store.s3.amazonaws.com/0/245792/0218f4ba-b158-7f8e-3398-2b1b0d73f52b.jpeg)
```python:detect_image.py
API_KEY = 'aaabbbccc'
API_SECRET = 'xxxyyyzzz'
```

3. Put your image under `./image/original`

```
./image/original/
├── person0
│   ├── person0_0.jpg
│   ├── person0_1.jpg
│   ├── person0_2.jpg
│   ├── person0_3.jpg
│   └── person0_4.jpg
├── person1
│   ├── person1_0.jpg
│   ├── person1_1.jpg
│   ├── person1_2.jpg
│   ├── person1_3.jpg
│   └── person1_4.jpg
└── person2
    ├── person2_0.jpg
    ├── person2_1.jpg
    ├── person2_2.jpg
    ├── person2_3.jpg
    └── person2_4.jpg
```
4. Face++ Detect API にリクエストを送り, レスポンスをJSON形式で保存します
```
$ python detect_image.py
```

5. 保存した JSONファイルをロードして, 取得した数値から顔を切り取ります
```
$ python clip_face.py
```

## Author
[Doarakko](https://github.com/Doarakko)

## License
MIT
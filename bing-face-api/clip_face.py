import bing_face_api as bfa

if __name__ == '__main__':
    '''
    コマンドライン引数を使用する場合
    顔認識する画像のディレクトリ
    search_dir = sys.argv[0]
    '''
    # 顔認識する画像のディレクトリ
    search_dir = "./image/original/"
    # 顔認識する画像のファイル名を取得
    img_path_list = bfa.get_image_path_list(search_dir)
    # 顔認識
    bfa.detect_image(img_path_list)

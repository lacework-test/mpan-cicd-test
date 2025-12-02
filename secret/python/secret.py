def main():
    # SECRET: Adobe Credentials
    # (?i)adobe[a-z0-9_]*(?:[\\:\\s]+)?(?:string|String|str)?['\\\"]?\\s*(=|>|:=|\\|\\|:|<=|=>|:)\\s*['\\\"]([A-Za-z0-9]{32,40})['\\\"]
    adobe_key = "AB12CD34EF56GH78JK90LM12NO34PQ56"

    # SECRET: Adobe Credentials
    # (?i)adobe[a-z0-9_]*(?:[\\:\\s]+)?(?:string|String|str)?['\\\"]?\\s*(=|>|:=|\\|\\|:|<=|=>|:)\\s*['\\\"](p8e-)(?i)[a-z0-9\\-_=+]{28,64}['\\\"]
    adobe_secret = "p8e-9z8y7x6w5v4u3t2s1r0_qp-on+ml_kj"

    # SECRET: Adobe Credentials
    # (=|>|:=|\\|\\|:|<=|=>|:)\\s*(p8e-)(?i)[a-z0-9\\-_=+]{28,64}
    token = p8e-9z8y7x6w5v4u3t2s1r0_qp-on+ml_kj



    # SECRET: Age Secret Key Credentials
    # ['\"]AGE-SECRET-KEY-[A-Za-z0-9]{54,}['\"]
    age_secret = "AGE-SECRET-KEY-abcdefghijABCDEFGHIJ1234567890klmnopqrstKLMNOPQRST1234567890"

    # SECRET: Age Secret Key Credentials
    # (=|>|:=|\|\|:|<=|=>|:)\s*AGE-SECRET-KEY-[A-Za-z0-9]{54,}
    token = AGE-SECRET-KEY-abcdefghijABCDEFGHIJ1234567890klmnopqrstKLMNOPQRST1234567890


    # SECRET: Alibaba Credentials
    # ['\"]LTAI[A-Za-z0-9]{12,32}['\"]
    alibaba_key = "LTAIabcdefghij1234567890klmnop"

    # SECRET: Alibaba Credentials
    # (=|>|:=|\|\|:|<=|=>|:)\s*LTAI[A-Za-z0-9]{12,32}
    alibaba_key = LTAIabcdefghij1234567890klmnop

    # SECRET: Alibaba Credentials
    # (?i)alibaba[a-z0-9_]*(?:[\\:\\s]+)?(?:string|String|str)?['\\\"]?\s*(=|>|:=|\\|\\|:|<=|=>|:)\s*['\\\"]([A-Za-z0-9+/=]{30,50})['\\\"]
    alibaba_secret = "alibabaABCDEFGHIJ1234567890klmnopqrstuv+/"


    # SECRET: Asana Credentials
    # (?i)asana[a-z0-9_ .\\-,]*(?:[\\:\\s]+)?(?:string|String|str)?['\\\"]?\s*(=|>|:=|\|\|:|<=|=>|:)\s*['\\\"]([0-9]{16})['\\\"]
    asana_id = "1234567890123456"

    # SECRET: Asana Credentials
    # (?i)asana[a-z0-9_ .\\-,]*(?:[\\:\\s]+)?(?:string|String|str)?['\\\"]?\s*(=|>|:=|\|\|:|<=|=>|:)\s*['\\\"]([a-z0-9]{32})['\\\"]
    asana_secret = "abcdefghij1234567890klmnopqrstuv"


    # SECRET: Asymmetric Encryption Credentials
    open_ssh_private_key: str =
    -----BEGIN OPENSSH PRIVATE KEY-----
    b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtzc2gtZW
    QyNTUxOQAAACAQ2rw9FnMxDHndk1CB1r3j17FMGB1sDhSXueP1oJOHPgAAAKBFZoirRWaI
    qwAAAAtzc2gtZWQyNTUxOQAAACAQ2rw9FnMxDHndk1CB1r3j17FMGB1sDhSXueP1oJOHPg
    AAAECtwceU5L0FZZliRWJo+5for6JS60fevlB9XmLf4XJTLBDavD0WczEMed2TUIHWvePX
    sUwYHWwOFJe54/Wgk4c+AAAAFnRvbnl0b255d3VAaG90bWFpbC5jb20BAgMEBQYH
    -----END OPENSSH PRIVATE KEY-----

    # SECRET: Asymmetric Encryption Credentials
    private_key: str =
    -----BEGIN PRIVATE KEY-----
    MIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQ
    C5dZm6rVfRWcb0Ve6ZrNo1GL/QBLXqExnpvng9Pv9gzKdcfGvF
    GqHTHBuDqfRHAdBmoB+Sp9S8XkUsHBe7Pim2448quBPPK15Ldh
    8fvQpNNHPE+0rFNbVOX79YSx8MmeEx2jQTS/wpEbFP0E1mCZBo
    UCHyhh4DlrDptHGYMtG40RQp3qupNkD9l5r72J2AT2/434xFJY
    -----END PRIVATE KEY-----

    # SECRET: Atlassian Credentials
    # Atlassian Generic (24–192)
    # (?i)atlassian[a-z0-9_]*(?:[\\:\\s]+)?(?:string|String|str)?['\\\"]?\\s*(=|>|:=|\\|\\|:|<=|=>|:)\\s*['\\\"]([a-zA-Z0-9+_/=\\-]{24,192})['\\\"]
    atlassian_key = "abcdefghijABCDEFGHIJ1234567890+/=--__1234567890"


    # Aws Credentials (bugs: 5682, 5683)
    # Aws Credentials (bugs: 5682, 5683)
    # Aws Credentials (bugs: 5682, 5683)


    # Aws String Credentials (bug 5683)
    # Aws String Credentials (bug 5683)
    # Aws String Credentials (bug 5683)


    # SECRET: Azure Credentials
    # (?i)azure[a-z0-9_ .\\-,]*(?:[\\:\\s]+)?(?:string|String|str)?['\\\"]?\\s*(=|>|:=|\\|\\|:|<=|=>|:)\\s*['\\\"]([A-Za-z0-9+/=]{60,90})['\\\"]
    azure_secret: str = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"


    # SECRET: Bcrypt Hash Credentials
    # (?:^|[^A-Za-z0-9./])\$2[abxy]?\$(0[4-9]|[12][0-9]|3[0-1])?\$?[A-Za-z0-9./]{45,60}(?:$|[^A-Za-z0-9./])
    bcrypt_hash: str = "$2b$12$abcdefghijklmnopqrstuvABCDEFGHIJKLMNOPQRSTUVWXYZabcde"


    # SECRET: Beamer Credentials
    # (?i)beamer[a-z0-9_ .\-,]*(?:[\:\s]+)?(?:string|String|str)?['\"]?\s*(=|>|:=|\|\|:|<=|=>|:)\s*['\"](b_[a-zA-Z0-9=_\-/+]{40,48})['\"]
    beamer_token: str = "b_ABCdef1234567890ghiJKLmnopQRSTUVwxyz+/-="


    # SECRET: Bitbucket Credentials
    # Short key (24–36 chars)
    bitbucket_key = "abcdefghijklmnopqrstuvwx12345678"

    # SECRET: Bitbucket Credentials
    # Long key (50–70 chars)
    bitbucket_token = "AbCdE-1234567890_qwertyuiopasdfghjklzxcvbnmMNBVCXZ"


    # SECRET: Clojars Credentials
    # (?i)(=|>|:=|\|\|:|<=|=>|:)\s*CLOJARS_[a-z0-9]{54,64}
    clojars_token = CLOJARS_abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz1234

    # SECRET: Clojars Credentials
    # (?i)['\"]CLOJARS_[a-z0-9]{54,64}['\"]
    clojars_token_quoted = "CLOJARS_abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz1234"


    # SECRET: Contentful Credentials
    # (?i)contentful[a-z0-9_]*(?:[\:\s]+)?(?:string|String|str)?['\"]?\s*(=|>|:=|\|\|:|<=|=>|:)\s*['\"]([a-z0-9-=_]{40,60})['\"]
    contentful_token: str = "contentful_abcdefghijklmnopqrstuvwxyz1234567890-=_"


    # SECRET: Databricks Credentials
    # ['\"]dapi[a-h0-9]{24,36}['\"]
    databricks_token: str = "dapi0123abcd4567efgh89a0bc1def2345678"

    # SECRET: Databricks Credentials
    # (=|>|:=|\|\|:|<=|=>|:)\s*dapi[a-h0-9]{24,36}
    databricks_key: str = "dapiabcdef0123456789abcdef0123456789"


    # SECRET: Digital Ocean Credentials
    # (?i)digital[a-z0-9_]*ocean[a-z0-9_ .\-,]*(?:[\:\s]+)?(?:string|String|str)?['\"]?\s*(=|>|:=|\|\|:|<=|=>|:)\s*['\"]([a-f0-9]{60,68})['\"]
    digital_ocean_token: str = "abcdef123456abcdef123456abcdef123456abcdef123456abcdef123456abcdef12"


    # SECRET: Discord Credentials
    # (?i)discord[a-z0-9_]*(?:[\\:\\s]+)?(?:string|String|str)?['\\\"]?\\s*(=|>|:=|\\|\\|:|<=|=>|:)\\s*['\\\"]([0-9]{14,32})['\\\"]
    discord_user_id: str = "123456789012345678"

    # SECRET: Discord Credentials
    # (?i)discord[a-z0-9_]*(?:[\\:\\s]+)?(?:string|String|str)?['\\\"]?\\s*(=|>|:=|\\|\\|:|<=|=>|:)\\s*['\\\"]([a-z0-9=_\\-\\.]{24,98})['\\\"]
    discord_token: str = "abcd1234_efgh5678-ijkl.mnopqrstu"


    # SECRET: Docker Authentication Credentials
    docker_auth_config: str =
    {
      "auths": {
        "https://index.docker.io/v1/": {
          "username": "yuri",
          "auth": "eXVyaTpzZWNyZXQ="
        }
      }
    }


    # SECRET: Doppler Credentials
    # ['\"](dp\.pt\.)(?i)[a-z0-9]{36,48}['\"]
    doppler_token: str = "dp.pt.abcd1234efgh5678ijkl9012mnop3456qrstuvwx"

    # SECRET: Doppler Credentials
    # (=|>|:=|\|\|:|<=|=>|:)\s*(dp\.pt\.)(?i)[a-z0-9]{36,48}
    dopplerKey: str = "dp.pt.zxcv0987bnml6543asdf2109qwer8765tyuiopgh"


    # SECRET: Dropbox Credentials
    # (?i)dropbox[a-z0-9_ .\-,]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(=|>|:=|\|:|<=|=>|:)\s*['"]([a-z0-9]{12,16})['"]
    dropbox_api_secret = "abcdefghijklmnop"

    # SECRET: Dropbox Credentials
    # (?i)dropbox[a-z0-9_ .\-,]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(=|>|:=|\|:|<=|=>|:)\s*['"](sl\.[a-z0-9\-=_]{124,148})['"]
    dropbox_short_lived_token = "sl.ABCD1234efgh5678ijkl9012mnop3456qrstuvwxyzABCD1234efgh5678ijkl9012mnop3456qrstuvwxyzABCD1234efgh5678ijkl9012mnop3456qrstuvwxyz"

    # SECRET: Dropbox Credentials
    # (?i)dropbox[a-z0-9_ .\-,]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(=|>|:=|\|:|<=|=>|:)\s*['"][a-z0-9]{11}(AAAAAAAAAA)[a-z0-9\-_=]{40,48}['"]
    dropbox = "1234567890aAAAAAAAAAA12345678901234567890abcdefghijabcdefghij"


    # SECRET: Duffel Credentials
    # ['"](duffel_(test|live)_(?i)[a-z0-9_-]{40,48})['"]
    duffel_token: str = "duffel_test_abcdefghij1234567890klmnopqrstuvwxyz1234567890"

    # SECRET: Duffel Credentials
    # (=|>|:=|\|:|<=|=>|:)\s*duffel_(test|live)_(?i)[a-z0-9_-]{40,48}
    duffel_key: str = "duffel_live_zxcv0987bnml6543asdf2109qwer8765tyuiop123456"


    # Dynatrace Credentials (bug 5683)
    # Dynatrace Credentials (bug 5683)
    # Dynatrace Credentials (bug 5683)
    # Dynatrace Credentials (bug 5683)
    # Dynatrace Credentials (bug 5683)

    # SECRET: Easypost Credentials
    # (?i)['"]EZ[APT]K[a-z0-9]{32,56}['"]
    easy_post_token: str = "EZTKabcdefghij1234567890klmnopqrstuv1234567890abcdef"

    # SECRET: Easypost Credentials
    # (=|>|:=|\|:|<=|=>|:)\s*EZ[APT]K[a-z0-9]{32,56}
    easy_post_key: str = "EZPKzxcv0987bnml6543asdf2109qwer8765tyuiop123456"


    # SECRET: Facebook Credentials
    # ((?i)(facebook[a-z0-9_]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(=|>|:=|\|:|<=|=>|:)\s*['"]([A-Za-z0-9]{24,48})['"]))
    facebook_token: str = "abcd1234efgh5678ijkl9012mnop3456qrst7890uvwx"


    # SECRET: Fastly Credentials
    # (?i)fastly[a-z0-9_ .\-,]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(=|>|:=|\|:|<=|=>|:)\s*['"]([A-Za-z0-9_\-\.+!=]{20,36})['"]
    fastly_token: str = "abcd1234efgh5678ijkl9012mnop3456qrst"


    # SECRET: Filezilla Credentials
    # <Pass(?: encoding=\"base64\")?>[^<]+</Pass>
    filezilla_password = '<Pass encoding="base64">U29tZUJhc2U2NERhdGE=</Pass>'


    # SECRET: Finicity Credentials
    # (?i)finicity[a-z0-9_ .\-,]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(=|>|:=|\|:|<=|=>|:)\s*['"]([a-z0-9]{20})['"]
    finicity_token: str = "1234567890acbdefghss"

    # SECRET: Finicity Credentials
    # (?i)finicity[a-z0-9_ .\-,]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(=|>|:=|\|:|<=|=>|:)\s*['"]([a-zA-Z0-9_\-+!]{80,120})['"]
    finicity_key: str = "abcd1234efgh5678ijkl9012mnop3456qrst7890uvwxyzABCD1234EFGH5678IJKL9012MNOP3456QRST7890"


    # SECRET: Flutterwave Credentials
    # ['"](FLW(PUB|SEC)K_TEST-(?i)[a-h0-9]{32}-X)['\"']
    flutterwave_secret_1 = "FLWSECK_TEST-abcdeabcde1234567890abcdeabcde12-X"

    # SECRET: Flutterwave Credentials
    # (=|>|:=|\\|\\|:|<=|=>|:)\s*(FLW(PUB|SEC)K_TEST-(?i)[a-h0-9]{32}-X)
    flutterwave_secret_2 = "FLWPUBK_TEST-abcdeabcde1234567890abcdeabcde12-X"

    # SECRET: Flutterwave Credentials
    # ['"](FLWSECK_TEST[a-h0-9]{12})["']
    flutterwave_secret_3 = "FLWSECK_TESTabcdeabcde12"

    # SECRET: Flutterwave Credentials
    # (=|>|:=|\\|\\|:|<=|=>|:)\s*(FLWSECK_TEST[a-h0-9]{12})
    flutterwave_secret_4 = "FLWSECK_TESTabcdeabcde12"


    # SECRET: Frame Io Credentials
    # ['"]fio-u-(?i)[a-z0-9-_=]{64}["']
    frameio_secret_1 = "fio-u-abcdefghij1234567890ABCDEFGHIJabcdefghij1234567890ABCDEFGHIJ1234"

    # SECRET: Frame Io Credentials
    # (=|>|:=|\\|\\|:|<=|=>|:)\s*fio-u-(?i)[a-z0-9-_=]{64}
    frameio_secret_2 = fio-u-abcdefghij1234567890ABCDEFGHIJabcdefghij1234567890ABCDEFGHIJ1234


    # SECRET: Gcp Credentials
    # (?s){[^}]*[\\\"]?type[\\\"]?\\s*:\\s*[\\\"]?service_account[\\\"]?[\\s\\S.]*?(private_key_id['\\\"]?\\s*\\:\\s*['\\\"]?[a-zA-Z0-9_$@\\-]+['\\\"]?)[^}]*?}
    gcp_service_account_key = {
        "type": "service_account",
        "project_id": "my-project",
        "private_key_id": "a1b2c3d4e5f6g7h8i9j0k1l2m3n4o1p2q3r4s5t6"
    }


    # SECRET: Generic Password Credentials
    # (?i)(password|passwd|user_pwd)[a-z0-9_]*(?:[\:\s]+)?(?:string|String|str)?['"]{0,2}\s*(=|>|:=|\|:|<=|=>|:)\s*[']{1,2}([^\s\{}?%[\]]+)[']{1,2}
    password_token: str = 'abcd1234efgh5678'

    # SECRET: Generic Password Credentials
    # (?i)(password|passwd|user_pwd)[a-z0-9_]*(?:[\:\s]+)?(?:string|String|str)?['"]{0,2}\s*(=|>|:=|\|:|<=|=>|:)\s*["]{1,2}([^\s\{}?%[\]]+)["]{1,2}
    passwd_key: str = "zxcv0987bnml6543"

    # Exclude
    # (?i)(password|passwd|user_pwd)[a-z0-9_]*['\"]{0,2}\s*(=|>|:=|\||: |<=|=>|:)\s*((True|true|false|False|0|1)|['\"](True|true|false|False|0|1)['\"]|(['\"]{1,2}([^\s\{\}\?\%\[\]]+)['\"]{1,2}[\.]))
    password = "true"


    # SECRET: Generic Secret Key Credentials
    # (?i)['"]?\bsecret[a-z0-9_]*key[a-z0-9_]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(:)\s*([A-Za-z0-9/+=]{16,128})
    secret_key_token: str = "abcd1234efgh5678ijkl9012"

    # SECRET: Generic Secret Key Credentials
    # (?i)['"]?\bsecret[a-z0-9_]*key[a-z0-9_]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(=|:=|=>|:)\s*['"]([a-zA-Z0-9_\-\.\$@*\^+!\[\]\(\)^=]{16,128})['"]
    secret_key_quoted: str = "zxcv0987bnml6543.asdf@123"

    # Excluded by: (?i)['"]?\bsecret[a-z0-9_]*key[a-z0-9_]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(=|:=|=>|:)\s*['"]\$([a-zA-Z0-9_\-\.\$@*\^+!\[\]\(\)^=]{15,128})['"]
    secret_key_excluded: str = "$zxcv0987bnml6543.asdf@123"


    # SECRET: Github Credentials
    # ['"]gh[pousr]_[0-9a-zA-Z]{32,40}['"]
    github_token = "ghp_abcdefghijABCDEFGHIJabcdefghij123456"

    # SECRET: Github Credentials
    # (=|>|:=|\||: |<=|=>|:)\s*gh[pousr]_[0-9a-zA-Z]{32,40}
    github_token_raw = "ghs_abcdefghijABCDEFGHIJabcdefghij123456"

    # SECRET: Github Credentials
    # ['"]?github_pat_[0-9a-zA-Z_\-]{60,}['"]?
    github_pat = "github_pat_abcdefghijABCDEFGHIJ1234567890abcdefghijABCDEFGHIJ1234567890abcdefghijABCDEFGHIJ12"


    # SECRET: Gitlab Credentials
    # ['"]gl(pat|cict)-[A-Za-z0-9_-]{20,}['"]
    gitlab_token = "glpat-abcdefghijABCDEFGHIJabcdefghij123456"

    # SECRET: Gitlab Credentials
    # (=|>|:=|\||: |<=|=>|:)\s*gl(pat|cict)-[A-Za-z0-9_-]{20,}
    gitlab_token_raw = "glcict-abcdefghijABCDEFGHIJabcdefghij123456"


    # SECRET: Go Cardless Credentials
    # ['"]live_(?i)[a-z0-9\-_=]{40}['"]
    gocardless_token = "live_abcdefghijABCDEFGHIJabcdefghij1234567890"

    # SECRET: Go Cardless Credentials
    # (=|>|:=|\||: |<=|=>|:)\s*live_(?i)[a-z0-9\-_=]{40}
    gocardless_token_raw = "live_abcdefghijABCDEFGHIJabcdefghij1234567890"


    # SECRET: Grafana Credentials
    # (?i)grafana[a-z0-9_]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(=|>|:=|\|:|<=|=>|:)\s*['"]?eyJrIjoi(?i)[a-z0-9\-_=]{72,126}['"]?
    grafana_token: str = "eyJrIjoiabcd1234efgh5678ijkl9012mnop3456qrst7890uvwxyz1234567890abcdef1234567890abcdefghij"

    # SECRET: Grafana Credentials
    # (?i)grafana[a-z0-9_]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(=|>|:=|\|:|<=|=>|:)\s*['"]?glsa_[a-zA-Z0-9_\-]{36,}['"]?
    grafana_key: str = "glsa_abcd1234efgh5678ijkl9012mnop3456qrst7890"


    # SECRET: Hashicorp Credentials
    # ['"](?i)[a-z0-9]{10,24}.atlasv1.[a-z0-9-_=]{60,70}["']
    hashicorp_secret_1 = "abcdefghij.atlasv1.1234567890ABCDEFGHIJabcdefghij1234567890ABCDEFGHIJabcdefghij"

    # SECRET: Hashicorp Credentials
    # (=|>|:=|\\|\\|:|<=|=>|:)\s*(?i)[a-z0-9]{10,24}\\.atlasv1\\.[a-z0-9\\-_=]{60,70}
    hashicorp_secret_2 = abcdefghij.atlasv1.1234567890ABCDEFGHIJabcdefghij1234567890ABCDEFGHIJabcdefghij

    # SECRET: Hashicorp Credentials
    # ['"]tfp_[A-Za-z0-9]{20,40}["']
    hashicorp_secret_3 = "tfp_abcdefghij1234567890ABCD"

    # SECRET: Hashicorp Credentials
    # (=|>|:=|\\|\\|:|<=|=>|:)\s*tfp_[A-Za-z0-9]{20,40}
    hashicorp_secret_4 : tfp_abcdefghij1234567890ABCD


    # SECRET: Heroku Credentials
    # (?i)heroku[a-z0-9_\]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(=|>|:=|\|:|<=|=>|:)\s*['"]?([0-9A-F]{8}-[0-9A-F]{4}-[0-9A-F]{4}-[0-9A-F]{4}-[0-9A-F]{12})['"]?
    heroku_token: str = "12345678-abcd-1234-abcd-1234567890ab"

    # SECRET: Heroku Credentials
    # (?i)heroku[a-z0-9_\]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(=|>|:=|\|:|<=|=>|:)\s*['"]?(HRKU-[a-zA-Z0-9\_\-=+]{48,72})['"]?
    heroku_hrk_token: str = "HRKU-abcd1234efgh5678ijkl9012mnop3456qrst7890uvwxyz1234567890abcd"

    # SECRET: Heroku Credentials
    # (?i)heroku[a-z0-9_\]*key[a-z0-9_\]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(=|>|:=|\|:|<=|=>|:)\s*['"]?([A-Za-z0-9/+=]{16,48})['"]?
    heroku_key: str = "abcd1234efgh5678ijkl9012"


    # SECRET: Htpasswd Credentials
    # ^[a-zA-Z0-9_-]+:\$apr1\$[A-Za-z0-9./]{8}\$[A-Za-z0-9./]{22}$
    htpasswd_value: str = "user1:$apr1$abcdefgh$1234567890123456789012"

    # SECRET: Htpasswd Credentials
    # (?i)[a-z0-9_]*htpasswd[a-z0-9_]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(=|>|:=|\|:|<=|=>|:)\s*['"]?([a-zA-Z0-9_-]+:\$apr1\$[A-Za-z0-9./]{8}\$[A-Za-z0-9./]{22})['"]?
    user_htpasswd: str = "admin:$apr1$ijklmnop$0987654321098765432109"


    # SECRET: Http Auth Header Credentials
    # (?i)Authorization[\\]?['"]?\s*(=|:=|:|=>)\s*[\\]?['"]?Basic\s+[A-Za-z0-9+/]{16,}={0,2}['"]?
    auth_header_basic: str = "Authorization: Basic YWRtaW46c2VjdXJlcGFzczEyMw=="

    # SECRET: Http Auth Header Credentials
    # (?i)Authorization[\\]?['"]?\s*(=|:=|:|=>)\s*[\\]?['"]?Bearer\s+([A-Za-z0-9-._~+/]{20,})\s*['"]?
    auth_header_bearer: str = "Authorization: Bearer abcdefghijklmnopqrst1234567890"


    # SECRET: Hubspot Credentials
    # (?i)hubspot[a-z0-9_]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(=|>|:=|\|:|<=|=>|:)\s*['"]([a-z0-9]{8}-[a-z0-9]{4}-[a-z0-9]{4}-[a-z0-9]{4}-[a-z0-9]{12})['"]
    hubspot_token: str = "abcd1234-efgh-5678-ijkl-9012mnop3456"

    # SECRET: Hubspot Credentials
    # (?i)hubspot[a-z0-9_]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(=|>|:=|\|:|<=|=>|:)\s*['"]pat-[a-z0-9]{1,4}-[a-z0-9]{8}-[a-z0-9]{4}-[a-z0-9]{4}-[a-z0-9]{4}-[a-z0-9]{12}['"]
    hubspot_pat: str = "pat-ab12-abcd1234-efgh-5678-ijkl-9012mnop3456"


    # SECRET: Hugging Face Token Credentials
    # ['"](hf_|api_org_|hf_app_)[a-zA-Z0-9]{30,64}['"]
    hugging_face_token: str = "hf_abcd1234efgh5678ijkl9012mnop3456"

    # SECRET: Hugging Face Token Credentials
    # (=|>|:=|\|:|<=|=>|:)\s*(hf_|api_org_|hf_app_)[a-zA-Z0-9]{30,64}
    hugging_face_key: str = api_org_1234567890abcdefeior1234567890ab


    # SECRET: Intercom Credentials
    # (?i)intercom[a-z0-9_ .,-]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(=|>|:=|\|:|<=|=>|:)\s*['"]([A-Za-z0-9=\-_]{48,72})['"]
    intercom_token: str = "abcd1234efgh5678ijkl9012mnop3456qrst7890uvwxyz1234567890abcdefghij"

    # SECRET: Intercom Credentials
    # (?i)intercom[a-z0-9_ .,-]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(=|>|:=|\|:|<=|=>|:)\s*['"]([a-h0-9]{8}-[a-h0-9]{4}-[a-h0-9]{4}-[a-h0-9]{4}-[a-h0-9]{12})['"]
    intercom_uuid: str = "abcd1234-efgh-5678-abcd-9012efgh3456"


    # SECRET: Ionic Credentials
    # (?i)ionic[a-z0-9_ .,-]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(=|>|:=|\|:|<=|=>|:)\s*['"](ion_[A-Za-z0-9]{32,45})['"]
    ionic_token: str = "ion_abcd1234efgh5678ijkl9012mnop3456qrst"


    # SECRET: Jekyll Token Credentials
    # (?i)[a-z0-9_]*jekyll[a-z0-9_]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(=|>|:=|\|:|<=|=>|:)\s*[a-zA-Z0-9]{24,}
    jekyll_token: str = abcd1234efgh5678ijkl9012mnop3456

    # SECRET: Jekyll Token Credentials
    # (?i)[a-z0-9_]*jekyll[a-z0-9_]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(=|>|:=|\|:|<=|=>|:)\s*['"][a-zA-Z0-9]{24,}['"]
    jekyll_key: str = "1234567890abcdef1234567890abcdef"


    # SECRET: Jwt Token Credentials
    # \b(?i)jwt[a-z0-9_]*['"]?\s*(=|:|=>|:=)\s*['"]([A-Za-z0-9-_]+\.[A-Za-z0-9-_]+\.[A-Za-z0-9-_]+)['"]
    jwt_token: str = "eyJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJ1c2VyIn0.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"

    # SECRET: Jwt Token Credentials
    # eyJ[a-zA-Z0-9_\-]+\.(eyJ[A-Za-z0-9-_]+)\.[A-Za-z0-9-_]+
    jwt_auth: str = "eyJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJ1c2VyIn0.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"

    # SECRET: Jwt Token Credentials
    # eyJ[a-zA-Z0-9_\-]+\.[A-Za-z0-9-_]+\.[A-Za-z0-9-_]+\.[A-Za-z0-9-_]+\.[A-Za-z0-9-_]+
    jwt_extended: str = "eyJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJ1c2VyIn0.abcd1234.efgh5678.ijkl9012"


    # SECRET: Launchdarkly Token Credentials
    # ['"]api-[a-z0-9]{8}-[a-z0-9]{4}-[a-z0-9]{4}-[a-z0-9]{4}-[a-z0-9]{12}['"]
    launch_darkly_token: str = "api-abcd1234-efgh-5678-ijkl-9012mnop3456"

    # SECRET: Launchdarkly Token Credentials
    # ['"]api-[a-f0-9]{32}['"]
    launch_darkly_key: str = "api-abcdef0123456789abcdef0123456789"

    # SECRET: Launchdarkly Token Credentials
    # (=|>|:=|\|:|<=|=>|:)\s*api-[a-z0-9]{8}-[a-z0-9]{4}-[a-z0-9]{4}-[a-z0-9]{4}-[a-z0-9]{12}
    launch_darkly_uuid: str = api-abcd1234-efgh-5678-ijkl-9012mnop3456

    # SECRET: Launchdarkly Token Credentials
    # (=|>|:=|\|:|<=|=>|:)\s*api-[a-f0-9]{32}
    launch_darkly_hex: str = api-1234567890abcdef1234567890abcdef


    # SECRET: Linear Credentials
    # ['"]lin_api_[A-Za-z0-9]{32,44}['"]
    linear_api_token: str = "lin_api_abcd1234efgh5678ijkl9012mnop3456"

    # SECRET: Linear Credentials
    # (=|>|:=|\|:|<=|=>|:)\s*lin_api_[A-Za-z0-9]{32,44}
    linear_api_key: str = lin_api_1234567890abcdef1234567890abcdef

    # SECRET: Linear Credentials
    # (?i)linear[a-z0-9_ .,-]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(=|>|:=|\|:|<=|=>|:)\s*['"]([a-f0-9]{28,44})['"]
    linear_token: str = "abcdef0123456789abcdef0123456789"


    # SECRET: Linkedin Credentials
    # (?i)linkedin[a-z0-9_]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(=|>|:=|\|:|<=|=>|:)\s*['"]((?i)[a-z0-9_\.=+!-]{14,64})['"]
    linkedin_token: str = "abcd1234efgh5678_ijkl9012.mnop=qr+st"


    # SECRET: Lob Credentials
    # (?i)lob[a-z0-9_ .,-]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(=|>|:=|\|:|<=|=>|:)\s*['"]((live|test)_[a-f0-9]{28,38})['"]
    lob_token: str = "live_abcdef0123456789abcdef0123456789"

    # SECRET: Lob Credentials
    # (?i)lob[a-z0-9_ .,-]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(=|>|:=|\|:|<=|=>|:)\s*['"]((live|test)_pub_[a-f0-9]{28,35})['"]
    lob_pub_token: str = "live_pub_1234567890abcdef1234567890ab"


    # SECRET: Mailchimp Token Credentials
    # (?i)[a-z0-9_]*mailchimp[a-z0-9_]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(=|>|:=|\|:|<=|=>|:)\s*[a-zA-Z0-9]{24,42}-us[0-9]{1,3}
    mailchimp_token: str = abcd1234efgh5678ijkl9012mnop3456-us12

    # SECRET: Mailchimp Token Credentials
    # (?i)[a-z0-9_]*mailchimp[a-z0-9_]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(=|>|:=|\|:|<=|=>|:)\s*['"][a-zA-Z0-9]{24,42}-us[0-9]{1,3}['"]
    mailchimp_key: str = "1234567890abcdef1234567890abcdef-us12"


    # SECRET: Mailgun Credentials
    # (?i)mailgun[a-z0-9_]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(=|>|:=|\|:|<=|=>|:)\s*['"]((pub)?key-[a-f0-9]{32})['"]
    mailgun_token: str = "key-abcdef0123456789abcdef0123456789"

    # SECRET: Mailgun Credentials
    # (?i)mailgun[a-z0-9_]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(=|>|:=|\|:|<=|=>|:)\s*['"]([a-h0-9]{32}-[a-h0-9]{8}-[a-h0-9]{8})['"]
    mailgun_key: str = "abcdef0123456789abcdef0123456789-abcdefgh-abcdefgh"


    # SECRET: Mapbox Credentials
    # (?i)(pk\.[a-z0-9\-_]{54,100}\.[a-z0-9\-_]{16,32})
    mapbox_token: str = "pk.abcd1234efgh5678ijkl9012mnop3456qrst7890uvwxyz1234567890ab.abcdefghijklmnop"


    # SECRET: Messagebird Credentials
    # (?i)messagebird[a-z0-9_]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(=|>|:=|\|:|<=|=>|:)\s*['"]((?i)[a-z0-9]{20,48})['"]
    messagebird_token: str = "abcd1234efgh5678ijkl9012mnop3456"

    # SECRET: Messagebird Credentials
    # (?i)messagebird[a-z0-9_]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(=|>|:=|\|:|<=|=>|:)\s*['"]((?i)[a-h0-9]{8}-[a-h0-9]{4}-[a-h0-9]{4}-[a-h0-9]{4}-[a-h0-9]{12})['"]
    messagebird_uuid: str = "abcd1234-efgh-5678-abcd-9012efgh3456"


    # SECRET: Mongodb Connection String Credentials
    # \bmongodb(?:\+srv)?:\/\/(?!<)[^\s:@\/]+:[^\s:@\/]+@[^\s]+
    mongo_db_connection: str = "mongodb://admin:securepass123@cluster0.mongodb.net"


    # SECRET: Mysql Connection String Credentials
    # (?i)`?(mysql|jdbc:mysql)://((?!{.*})[^\s:@/]+)(:((?!{.*})[^\s:@/]+))?@[^\\s:/]+(:[^\\s/]+)?/((?!{.*})[^\\s?]+)`?
    mysql_connection: str = "mysql://admin:securepass123@localhost:3306/mydb"

    # SECRET: Mysql Connection String Credentials
    # (?i)`?(mysql|jdbc:mysql)://[^\\s:/]+(:[^\\s/]+)?/((?!{.*})[^\\s?]+)\?.*?((user=((?!{.*})[^&\s]+))|(password=((?!{.*})[^&\s]+)))
    mysql_connection_params: str = "mysql://localhost:3306/mydb?user=admin&password=securepass123"


    # SECRET: New Relic Credentials
    # NRAK-[A-Z0-9]{16,30}\b
    new_relic_api_key: str = "NRAK-XYZ78901PQR23456STUV"

    # SECRET: New Relic Credentials
    # NRJS-[a-z0-9]{16,30}\b
    new_relic_js_key: str = "NRJS-xyzw7890pqrs1234tuv"

    # SECRET: New Relic Credentials
    # NRBR-[a-z0-9]{16,30}\b
    new_relic_browser_key: str = "NRBR-9012345678xyzab12345"

    # SECRET: New Relic Credentials
    # (?i)newrelic[a-z0-9_]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(=|>|:=|\|:|<=|=>|:)\s*['"]((?i)[a-z0-9\-_]{44,82})['"]
    newrelic_token: str = "abcdef0123-ghij4567-klmn8901-pqrs2345-tuvwxyz6789-abcde"


    # SECRET: Npm Access Credentials
    # ['"]npm_(?i)[a-z0-9]{30,42}['"]
    npm_token: str = "npm_abcdef0123456789ghijklmnop1234567890"


    # SECRET: Openai Token Credentials
    # ['"]sk-[a-zA-Z0-9]{25,50}['"]
    openai_token: str = "sk-abcd1234EFGH5678ijkl9012MNOP3456"

    # SECRET: Openai Token Credentials
    # (=|>|:=|\|:|<=|=>|:)\s*sk-[a-zA-Z0-9]{25,50}
    openai_key: str = sk-12345678abcdEFGH9012ijklMNOP3456


    # SECRET: P12 Credentials
    # \.p12\s*(=|:=|:|=>|)\s*['"]([a-zA-Z0-9\-_./]{24,})['"]
    p12_cert: str = "cert.p12=abcd1234-efgh5678_ijkl9012.mnop3456"

    # SECRET: P12 Credentials
    # \.p12\s*(=|:=|:|=>|)\s*([a-zA-Z0-9\-_./]{24,})
    p12_path: str = cert.p12=12345678-abcd9012_efgh3456.ijkl7890


    # SECRET: Paypal Credentials
    # (?i)paypal[a-z0-9_ .,-]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(=|>|:=|\|:|<=|=>|:)\s*['"]([a-zA-Z0-9\-_]{24,84})['"]
    paypal_token: str = "abcd1234EFGH5678_ijkl9012-MNOP3456"

    # SECRET: Paypal Credentials
    # ['"]A21A[A-Za-z0-9-_]{92}['"]
    paypal_api_key: str = "A21A90123456XYZW7890abcd1234EFGH5678ijkl9012MNOPqrst3456UVWXyzab7890CDEF1234ghij5678klmn90122222"

    # SECRET: Paypal Credentials
    # (=|>|:=|\|:|<=|=>|:)\s*A21A[A-Za-z0-9-_]{92}
    paypal_api_secret: str = A21A90123456XYZW7890abcd1234EFGH5678ijkl9012MNOPqrst3456UVWXyzab7890CDEF1234ghij5678klmn90122222


    # SECRET: Planetscale Credentials
    # ['"]pscale_pw_(?i)[a-z0-9\-_\.]{36,52}['"]
    planetscale_key: str = "pscale_pw_abcd1234-efgh5678_ijkl9012.mnop3456qrst"

    # SECRET: Planetscale Credentials
    # ['"]pscale_tkn_(?i)[a-z0-9\-_\.]{36,52}['"]
    planetscale_token: str = "pscale_tkn_12345678-abcd9012_efgh3456.ijkl7890mnop"


    # SECRET: Postgresql Connection String Credentials
    # (?i)`?(postgres|jdbc:postgresql|postgresql)\:\/\/[^\s\:@\/]+(\:[^{}\s\:@\/]+)@[^{}\s\:\/]+(\:[^{}\s\/]+)?\/[^{}\s?]+`?
    postgres_connection: str = "postgres://admin:secure123@localhost:5432/mydb"

    # SECRET: Postgresql Connection String Credentials
    # (?i)`?(postgres|jdbc:postgresql|postgresql)\:\/\/[^\s\:\/]+(\:[^\s\/]+)?\/((?!{.*})[^\s?]+)\?.*?((user=((?!{.*})[^&\s]+))|(password=((?!{.*})[^&\s]+)))
    postgres_connection_params: str = "postgres://localhost:5432/mydb?user=admin&key=secure123"


    # SECRET: Postman Api Credentials
    # PMAK-(?i)[a-f0-9]{24}-[a-f0-9]{34}\b
    postman_api_key = "PMAK-abcdefabcd12345678901234-1234567890abcdefabcd1234567890abcd"


    # SECRET: Pulumni Api Credentials
    # pul-[a-f0-9]{40}\\b
    pulumni_api_key = "pul-abcdeabcde1234567890abcdeabcde1234567890"


    # SECRET: Putty Key Credentials
    # (?si)PuTTY-User-Key-File-\\d+: \\S+\\s+((?![\"]).)*?Private-Lines: [0-9]+(?:\\s+[A-Za-z0-9+/=]+)+(?:\\s+Private-MAC: [a-f0-9]+)?
    putty_key = "PuTTY-User-Key-File-3: ssh-rsa\nEncryption: none\nComment: user@host\nPrivate-Lines: 2\nabcdefghij1234567890ABCDEFGHIJ1234567890==\nabcdefghij1234567890ABCDEFGHIJ1234567890==\nPrivate-MAC: 1234567890abcdef1234567890abcdef1234567890"

    # SECRET: Putty Key Credentials
    # ['\""](?si)PuTTY-User-Key-File-\\d+:[^'\""]*Private-Lines: [0-9]+[^'\""]*Private-MAC: [a-f0-9]+['\""]
    putty_key = "PuTTY-User-Key-File-3: ssh-rsa\nEncryption: none\nComment: user@host\nPrivate-Lines: 2\nabcdefghij1234567890ABCDEFGHIJ1234567890==\nabcdefghij1234567890ABCDEFGHIJ1234567890==\nPrivate-MAC: 1234567890abcdef1234567890abcdef1234567890"


    # SECRET: Pypi Credentials
    # ['\""]pypi-AgEIcHlwaS5vcmc[A-Za-z0-9_-]{24,}['\""]
    pypi_api_key = "pypi-AgEIcHlwaS5vcmcabcdefghij1234567890-_34"

    # SECRET: Pypi Credentials
    # (=|>|:=|\|\||:|<|<=|=>|:)\\s*pypi-AgEIcHlwaS5vcmc[A-Za-z0-9_-]{24,}
    pypi_api_key = pypi-AgEIcHlwaS5vcmcabcdefghij1234567890ABCD-_


    # SECRET: Rubygems Api Credentials
    # rubygems_[a-f0-9]{40,52}\\b
    rubygems_api_key = "rubygems_abcdeabcde1234567890abcdeabcde1234567890"


    # SECRET: Sendgrid Api Credentials
    # SG\.(?i)[a-z0-9_\\-\\.]{54,78}\\b
    sendgrid_api_key = "SG.abcdefghij1234567890abcdefghij1234567_-.abcdefghij1234"


    # SECRET: Sendinblue Api Credentials
    # xkeysib-[a-f0-9]{64}-(?i)[a-z0-9]{16}\\b
    sendinblue_api_key = "xkeysib-abcdeabcde1234567890abcdeabcde1234567890abcdeabcde12345678901234-1234567890abcdef"


    # SECRET: Shippo Api Credentials
    # shippo_(live|test)_[a-f0-9]{36,48}\b
    shippo_api_key: str = "shippo_live_aabbccdeef0123456789aabbccdeef0123456789"


    # SECRET: Shopify Credentials
    # ['\""]shp(ss|at|ca|pa)_[a-zA-Z0-9]{32}['\""]
    shopify_api_key_1 = "shpat_abcdefghij1234567890ABCDEFGHIJ12"

    # SECRET: Shopify Credentials
    # (=|>|:=|\\|\\|:|<=|=>|:)\\s*shp(ss|at|ca|pa)_[a-zA-Z0-9]{32}
    shopify_api_key_2 = shpss_abcdefghij1234567890ABCDEFGHIJ12


    # SECRET: Slack Credentials
    # https://hooks.slack.com/services/[A-Za-z0-9+/]{40,}
    slack_api_key_1 = "https://hooks.slack.com/services/abcdefghij1234567890ABCDEFGHIJ1234567890+/"

    # SECRET: Slack Credentials
    # ['\""]xox[baprs]-[0-9a-zA-Z\\-]{10,120}['\""]
    slack_api_key_2 = "xoxb-abcdefghij1234567890ABCDEFGHIJ1234567890"

    # SECRET: Slack Credentials
    # (=|>|:=|\\|\\|:|<=|=>|:)\\s*xox[baprs]-[0-9a-zA-Z\\-]{10,120}
    slack_api_key_3 = xoxa-abcdefghij1234567890ABCDEFGHIJ1234567890-


    # SECRET: Stripe Credentials
    # (?i)stripe[a-z0-9_]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(=|>|:=|\|:|<=|=>|:)\s*['"](?i)(sk|pk)_(live|test)_[a-z0-9]{10,120}['"]
    stripe_api_key1: str = "sk_live_aabbccdeef0123456789"

    # SECRET: Stripe Credentials
    # (?i)stripe[a-z0-9_]*\s*(=|>|:=|\|:|<=|=>|:)\s*(?i)(sk|pk)_(live|test)_[a-z0-9]{10,120}
    stripe_api_key2 = pk_test_0123456789aabbccdeef

    # SECRET: Stripe Credentials
    # ['"]sk_(live|test)_[A-Za-z0-9]{10,32}['"]
    stripe_api_key3: str = "sk_live_AABBCCDDEF0123456789"

    # SECRET: Stripe Credentials
    # (=|>|:=|\|:|<=|=>|:)\s*sk_(live|test)_[A-Za-z0-9]{10,32}
    stripe_api_key4: str = sk_test_AABBCCDDEF0123456789


    # SECRET: Twilio Credentials
    # ['"]SK[0-9a-fA-F]{32}['"]
    twilio_api_key1: str = "SKaabbccdeef0123456789abcdef01234567"

    # SECRET: Twilio Credentials
    # (=|>|:=|\|:|<=|=>|:)\s*SK[0-9a-fA-F]{32}
    twilio_api_key2: str = SKaabbccdeef0123456789abcdef01234567

    # SECRET: Twilio Credentials
    # (?i)twilio[a-z0-9_]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(=|>|:=|\|:|<=|=>|:)\s*['"]([A-Za-z0-9+/]{24,50})['"]
    twilio_api_key3: str = "aabbccdeef0123456789abcd"


    # SECRET: Twitch Credentials
    # (?i)twitch[a-z0-9_]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(=|>|:=|\|:|<=|=>|:)\s*['"](?i)[a-z0-9]{24,36}['"]
    twitch_api_key: str = "aabbccdeef0123456789aabbccdeef"


    # SECRET: Twitter Credentials
    # (?i)twitter[a-z0-9_]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(=|>|:=|\|:|<=|=>|:)\s*['"]([A-Za-z0-9+/\-_+]{24,64})['"]
    twitter_api_key: str = "aabbccdeef0123456789abcdABCD12"


    # SECRET: Typeform Api Credentials
    # (?i)typeform[a-z0-9_]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(=|>|:=|\|:|<=|=>|:)\s*['"](tfp_[a-z0-9\-_\.=]{54,64})['"]
    typeform_api_key: str = "tfp_aabbccdeef0123456789_.-=aabbccdeef0123456789_.-=aabbcc"


    # SECRET: Wordpress Auth Credentials
    # define(s*['"]*(DB_PASSWORD|MYSQL_PASSWORD|DB_PASS|DATABASE_PASSWORD)['"]*s*,s*['"][^'"]+['"]s*)
    wordpress_api_key_1 = "define('DB_PASSWORD', 'abcdefghij1234567890ABCDEFGHIJ12!@#$')"

    # SECRET: Wordpress Auth Credentials
    # define(s*['"]*(AUTH_KEY|SECURE_AUTH_KEY|LOGGED_IN_KEY|NONCE_KEY|AUTH_SALT|SECURE_AUTH_SALT|LOGGED_IN_SALT|NONCE_SALT)['"]*s*,s*['"][^'"]+['"]s*)
    wordpress_api_key_2 = "define('AUTH_KEY', 'abcdefghij1234567890ABCDEFGHIJ12!@#$')"


    # SECRET: Zoom Credentials
    # (?i)zoom[a-z0-9_ .\-,]*(?:[\:\s]+)?(?:string|String|str)?['\"]?\s*(=|>|:=|\|\|:|<=|=>|:)\s*['\"]([a-zA-Z0-9\-_]{24,60})['\"]
    zoom_token: str = "AbcdEFGH1234ijkl-5678MNOP_qrstUVWX90yz12"

if __name__ == "__main__":
    main()
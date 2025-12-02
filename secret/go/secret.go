package main

import "fmt"

func main() {
    // SECRET: Adobe Credentials
    // (?i)adobe[a-z0-9_]*(?:[\\:\\s]+)?(?:string|String|str)?['\\\"]?\\s*(=|>|:=|\\|\\|:|<=|=>|:)\\s*['\\\"]([A-Za-z0-9]{32,40})['\\\"]
    adobeKey := "AB12CD34EF56GH78JK90LM12NO34PQ56"

    // SECRET: Adobe Credentials
    // (?i)adobe[a-z0-9_]*(?:[\\:\\s]+)?(?:string|String|str)?['\\\"]?\\s*(=|>|:=|\\|\\|:|<=|=>|:)\\s*['\\\"](p8e-)(?i)[a-z0-9\\-_=+]{28,64}['\\\"]
    adobeSecret := "p8e-z1y2x3w4v5u6t7s8r9q0po_nm-lk+ji"

    // SECRET: Adobe Credentials
    // (=|>|:=|\\|\\|:|<=|=>|:)\\s*(p8e-)(?i)[a-z0-9\\-_=+]{28,64}
    secret := p8e-z1y2x3w4v5u6t7s8r9q0po_nm-lk+ji


    // SECRET: Age Secret Key Credentials
    // ['\"]AGE-SECRET-KEY-[A-Za-z0-9]{54,}['\"]
    ageSecret := "AGE-SECRET-KEY-abcdefghijABCDEFGHIJ1234567890klmnopqrstKLMNOPQRST1234567890"

    // SECRET: Age Secret Key Credentials
    // (=|>|:=|\|\|:|<=|=>|:)\s*AGE-SECRET-KEY-[A-Za-z0-9]{54,}
    secret := AGE-SECRET-KEY-abcdefghijABCDEFGHIJ1234567890klmnopqrstKLMNOPQRST1234567890


    // SECRET: Alibaba Credentials
    // ['\"]LTAI[A-Za-z0-9]{12,32}['\"]
    alibabaKey := "LTAIabcdefghij1234567890klmnop"

    // SECRET: Alibaba Credentials
    // (=|>|:=|\|\|:|<=|=>|:)\s*LTAI[A-Za-z0-9]{12,32}
    alibaba_key := LTAIabcdefghij1234567890klmnop

    // SECRET: Alibaba Credentials
    // (?i)alibaba[a-z0-9_]*(?:[\\:\\s]+)?(?:string|String|str)?['\\\"]?\s*(=|>|:=|\\|\\|:|<=|=>|:)\s*['\\\"]([A-Za-z0-9+/=]{30,50})['\\\"]
    alibaba_secret := "alibabaABCDEFGHIJ1234567890klmnopqrstuv+/"


    // SECRET: Asana Credentials
    // (?i)asana[a-z0-9_ .\\-,]*(?:[\\:\\s]+)?(?:string|String|str)?['\\\"]?\s*(=|>|:=|\|\|:|<=|=>|:)\s*['\\\"]([0-9]{16})['\\\"]
    asanaID := "1234567890123456"

    // SECRET: Asana Credentials
    // (?i)asana[a-z0-9_ .\\-,]*(?:[\\:\\s]+)?(?:string|String|str)?['\\\"]?\s*(=|>|:=|\|\|:|<=|=>|:)\s*['\\\"]([a-z0-9]{32})['\\\"]
    asanaSecret := "abcdefghij1234567890klmnopqrstuv"


    // SECRET: Asymmetric Encryption Credentials
    var openSshPrivateKey string =
    -----BEGIN OPENSSH PRIVATE KEY-----
    b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtzc2gtZW
    QyNTUxOQAAACAQ2rw9FnMxDHndk1CB1r3j17FMGB1sDhSXueP1oJOHPgAAAKBFZoirRWaI
    qwAAAAtzc2gtZWQyNTUxOQAAACAQ2rw9FnMxDHndk1CB1r3j17FMGB1sDhSXueP1oJOHPg
    AAAECtwceU5L0FZZliRWJo+5for6JS60fevlB9XmLf4XJTLBDavD0WczEMed2TUIHWvePX
    sUwYHWwOFJe54/Wgk4c+AAAAFnRvbnl0b255d3VAaG90bWFpbC5jb20BAgMEBQYH
    -----END OPENSSH PRIVATE KEY-----

    // SECRET: Asymmetric Encryption Credentials
    var privateKey string =
    -----BEGIN PRIVATE KEY-----
    MIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQ
    C5dZm6rVfRWcb0Ve6ZrNo1GL/QBLXqExnpvng9Pv9gzKdcfGvF
    GqHTHBuDqfRHAdBmoB+Sp9S8XkUsHBe7Pim2448quBPPK15Ldh
    8fvQpNNHPE+0rFNbVOX79YSx8MmeEx2jQTS/wpEbFP0E1mCZBo
    UCHyhh4DlrDptHGYMtG40RQp3qupNkD9l5r72J2AT2/434xFJY
    -----END PRIVATE KEY-----


    // SECRET: Atlassian Credentials
    // Atlassian Generic (24–192)
    // (?i)atlassian[a-z0-9_]*(?:[\\:\\s]+)?(?:string|String|str)?['\\\"]?\\s*(=|>|:=|\\|\\|:|<=|=>|:)\\s*['\\\"]([a-zA-Z0-9+_/=\\-]{24,192})['\\\"]
    atlassianKey := "abcdefghijABCDEFGHIJ1234567890+/=--__1234567890"


    // Aws Credentials (bugs: 5682, 5683)
    // Aws Credentials (bugs: 5682, 5683)
    // Aws Credentials (bugs: 5682, 5683)


    // Aws String Credentials (bug 5683)
    // Aws String Credentials (bug 5683)
    // Aws String Credentials (bug 5683)


    // SECRET: Azure Credentials
    // (?i)azure[a-z0-9_ .\\-,]*(?:[\\:\\s]+)?(?:string|String|str)?['\\\"]?\\s*(=|>|:=|\\|\\|:|<=|=>|:)\\s*['\\\"]([A-Za-z0-9+/=]{60,90})['\\\"]
    var azureSecret string = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"


    // SECRET: Bcrypt Hash Credentials
    // (?:^|[^A-Za-z0-9./])\$2[abxy]?\$(0[4-9]|[12][0-9]|3[0-1])?\$?[A-Za-z0-9./]{45,60}(?:$|[^A-Za-z0-9./])
    var bcryptHash string = "$2b$12$abcdefghijklmnopqrstuvABCDEFGHIJKLMNOPQRSTUVWXYZabcde"


    // SECRET: Beamer Credentials
    // (?i)beamer[a-z0-9_ .\-,]*(?:[\:\s]+)?(?:string|String|str)?['\"]?\s*(=|>|:=|\|\|:|<=|=>|:)\s*['\"](b_[a-zA-Z0-9=_\-/+]{40,48})['\"]
    var beamerToken string = "b_XYZabc1234567890+/ABCDEFGHIJKLMNOqrstuv="


    // SECRET: Bitbucket Credentials
    // (?i)bitbucket[a-z0-9_]*(?:[\:\s]+)?(?:string|String|str)?['\"]?\s*(=|>|:=|\|\|:|<=|=>|:)\s*['\"]([a-zA-Z0-9]{24,36})['\"]
    var bitbucket_key string = "Zyxwvutsrqponmlkjihgfedcba654321"

    // SECRET: Bitbucket Credentials
    // (?i)bitbucket[a-z0-9_]*(?:[\:\s]+)?(?:string|String|str)?['\"]?\s*(=|>|:=|\|\|:|<=|=>|:)\s*['\"]([a-zA-Z0-9_-]{50,70})['\"]
    var bitbucket_token string = "12345-ABCDE_fghijklmnopqrstuvwxyz0987654321QWERTYXY"


    // SECRET: Clojars Credentials
    // (?i)(=|>|:=|\|\|:|<=|=>|:)\s*CLOJARS_[a-z0-9]{54,64}
    var clojarsToken = CLOJARS_abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz1234

    // SECRET: Clojars Credentials
    // (?i)['\"]CLOJARS_[a-z0-9]{54,64}['\"]
    var clojarsTokenQuoted = "CLOJARS_abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz1234"


    // SECRET: Contentful Credentials
    // (?i)contentful[a-z0-9_]*(?:[\:\s]+)?(?:string|String|str)?['\"]?\s*(=|>|:=|\|\|:|<=|=>|:)\s*['\"]([a-z0-9-=_]{40,60})['\"]
    var contentful_token string = "contentful_abcdefghijklmnopqrstuvwxyz1234567890-=_"


    // SECRET: Databricks Credentials
    // ['\"]dapi[a-h0-9]{24,36}['\"]
    var databricks_token string = "dapi0123abcd4567efgh89a0bc1def2345678"

    // SECRET: Databricks Credentials
    // (=|>|:=|\|\|:|<=|=>|:)\s*dapi[a-h0-9]{24,36}
    var databricks_key string = "dapiabcdef0123456789abcdef0123456789"


    // SECRET: Digital Ocean Credentials
    // (?i)digital[a-z0-9_]*ocean[a-z0-9_ .\-,]*(?:[\:\s]+)?(?:string|String|str)?['\"]?\s*(=|>|:=|\|\|:|<=|=>|:)\s*['\"]([a-f0-9]{60,68})['\"]
    var digital_ocean_token string = "abcdef012345abcdef012345abcdef012345abcdef012345abcdef012345abcd"


    // SECRET: Discord Credentials
    // (?i)discord[a-z0-9_]*(?:[\\:\\s]+)?(?:string|String|str)?['\\\"]?\\s*(=|>|:=|\\|\\|:|<=|=>|:)\\s*['\\\"]([0-9]{14,32})['\\\"]
    var discord_user_id string = "123456789012345678"

    // SECRET: Discord Credentials
    // (?i)discord[a-z0-9_]*(?:[\\:\\s]+)?(?:string|String|str)?['\\\"]?\\s*(=|>|:=|\\|\\|:|<=|=>|:)\\s*['\\\"]([a-z0-9=_\\-\\.]{24,98})['\\\"]
    var discord_token string = "abcd1234_efgh5678-ijkl.mnopqrstu"


    // SECRET: Docker Authentication Credentials
    var dockerAuthConfig string =
    {
      "auths": {
        "https://index.docker.io/v1/": {
          "username": "yuri",
          "auth": "eXVyaTpzZWNyZXQ="
        }
      }
    }


    // SECRET: Doppler Credentials
    // ['\"](dp\.pt\.)(?i)[a-z0-9]{36,48}['\"]
    var doppler_token string = "dp.pt.abcd1234efgh5678ijkl9012mnop3456qrstuvwx"

    // SECRET: Doppler Credentials
    // (=|>|:=|\|\|:|<=|=>|:)\s*(dp\.pt\.)(?i)[a-z0-9]{36,48}
    var dopplerKey string = "dp.pt.zxcv0987bnml6543asdf2109qwer8765tyuiopgh"


    // SECRET: Dropbox Credentials
    // (?i)dropbox[a-z0-9_ .\-,]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(=|>|:=|\|:|<=|=>|:)\s*['"]([a-z0-9]{12,16})['"]
    dropboxApiSecret := "abcdefghijklmnop"

    // SECRET: Dropbox Credentials
    // (?i)dropbox[a-z0-9_ .\-,]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(=|>|:=|\|:|<=|=>|:)\s*['"](sl\.[a-z0-9\-=_]{124,148})['"]
    dropboxShortLivedToken := "sl.ABCD1234efgh5678ijkl9012mnop3456qrstuvwxyzABCD1234efgh5678ijkl9012mnop3456qrstuvwxyzABCD1234efgh5678ijkl9012mnop3456qrstuvwxyz"

    // SECRET: Dropbox Credentials
    // (?i)dropbox[a-z0-9_ .\-,]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(=|>|:=|\|:|<=|=>|:)\s*['"][a-z0-9]{11}(AAAAAAAAAA)[a-z0-9\-_=]{40,48}['"]
    dropbox := "1234567890aAAAAAAAAAA12345678901234567890abcdefghijabcdefghij"


    // SECRET: Duffel Credentials
    // ['"](duffel_(test|live)_(?i)[a-z0-9_-]{40,48})['"]
    var duffelToken string = "duffel_test_abcdefghij1234567890klmnopqrstuvwxyz1234567890"

    // SECRET: Duffel Credentials
    // (=|>|:=|\|:|<=|=>|:)\s*duffel_(test|live)_(?i)[a-z0-9_-]{40,48}
    var duffelKey string = "duffel_live_zxcv0987bnml6543asdf2109qwer8765tyuiop123456"


    // Dynatrace Credentials (bug 5683)
    // Dynatrace Credentials (bug 5683)
    // Dynatrace Credentials (bug 5683)
    // Dynatrace Credentials (bug 5683)
    // Dynatrace Credentials (bug 5683)

    // SECRET: Easypost Credentials
    // (?i)['"]EZ[APT]K[a-z0-9]{32,56}['"]
    var easyPostToken string = "EZTKabcdefghij1234567890klmnopqrstuv1234567890abcdef"

    // SECRET: Easypost Credentials
    // (=|>|:=|\|:|<=|=>|:)\s*EZ[APT]K[a-z0-9]{32,56}
    var easyPostKey string = "EZPKzxcv0987bnml6543asdf2109qwer8765tyuiop123456"


    // SECRET: Facebook Credentials
    // ((?i)(facebook[a-z0-9_]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(=|>|:=|\|:|<=|=>|:)\s*['"]([A-Za-z0-9]{24,48})['"]))
    var facebookToken string = "abcd1234efgh5678ijkl9012mnop3456qrst7890uvwx"


    // SECRET: Fastly Credentials
    // (?i)fastly[a-z0-9_ .\-,]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(=|>|:=|\|:|<=|=>|:)\s*['"]([A-Za-z0-9_\-\.+!=]{20,36})['"]
    var fastlyToken string = "abcd1234efgh5678ijkl9012mnop3456qrst"


    // SECRET: Filezilla Credentials
    // <Pass(?: encoding=\"base64\")?>[^<]+</Pass>
    var filezilla_password = '<Pass encoding="base64">U29tZUJhc2U2NERhdGE=</Pass>'


    // SECRET: Finicity Credentials
    // (?i)finicity[a-z0-9_ .\-,]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(=|>|:=|\|:|<=|=>|:)\s*['"]([a-z0-9]{20})['"]
    var finicityToken string = "1234567890acbdefghss"

    // SECRET: Finicity Credentials
    // (?i)finicity[a-z0-9_ .\-,]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(=|>|:=|\|:|<=|=>|:)\s*['"]([a-zA-Z0-9_\-+!]{80,120})['"]
    var finicityKey string = "abcd1234efgh5678ijkl9012mnop3456qrst7890uvwxyzABCD1234EFGH5678IJKL9012MNOP3456QRST7890"


    // SECRET: Flutterwave Credentials
    // ['"](FLW(PUB|SEC)K_TEST-(?i)[a-h0-9]{32}-X)['\"']
    const flutterwave_secret_1 = "FLWSECK_TEST-abcdeabcde1234567890abcdeabcde12-X";

    // SECRET: Flutterwave Credentials
    // (=|>|:=|\\|\\|:|<=|=>|:)\s*(FLW(PUB|SEC)K_TEST-(?i)[a-h0-9]{32}-X)
    const flutterwave_secret_2 = "FLWPUBK_TEST-abcdeabcde1234567890abcdeabcde12-X";

    // SECRET: Flutterwave Credentials
    // ['"](FLWSECK_TEST[a-h0-9]{12})["']
    const flutterwave_secret_3 = "FLWSECK_TESTabcdeabcde12";

    // SECRET: Flutterwave Credentials
    // (=|>|:=|\\|\\|:|<=|=>|:)\s*(FLWSECK_TEST[a-h0-9]{12})
    const flutterwave_secret_4 = "FLWSECK_TESTabcdeabcde12";


    // SECRET: Frame Io Credentials
    // ['"]fio-u-(?i)[a-z0-9-_=]{64}["']
    var frameio_secret_1 = "fio-u-abcdefghij1234567890ABCDEFGHIJabcdefghij1234567890ABCDEFGHIJ1234"

    // SECRET: Frame Io Credentials
    // (=|>|:=|\\|\\|:|<=|=>|:)\s*fio-u-(?i)[a-z0-9-_=]{64}
    var frameio_secret_2 = fio-u-abcdefghij1234567890ABCDEFGHIJabcdefghij1234567890ABCDEFGHIJ1234


    // SECRET: Gcp Credentials
    // (?s){[^}]*[\\\"]?type[\\\"]?\\s*:\\s*[\\\"]?service_account[\\\"]?[\\s\\S.]*?(private_key_id['\\\"]?\\s*\\:\\s*['\\\"]?[a-zA-Z0-9_$@\\-]+['\\\"]?)[^}]*?}
    var gcp_service_account_key = {
        "type": "service_account",
        "project_id": "my-project",
        "private_key_id": "a1b2c3d4e5f6g7h8i9j0k1l2m3n4o1p2q3r4s5t6"
    }


    // SECRET: Generic Password Credentials
    // (?i)(password|passwd|user_pwd)[a-z0-9_]*(?:[\:\s]+)?(?:string|String|str)?['"]{0,2}\s*(=|>|:=|\|:|<=|=>|:)\s*[']{1,2}([^\s\{}?%[\]]+)[']{1,2}
    var passwordToken string = 'abcd1234efgh5678'

    // SECRET: Generic Password Credentials
    // (?i)(password|passwd|user_pwd)[a-z0-9_]*(?:[\:\s]+)?(?:string|String|str)?['"]{0,2}\s*(=|>|:=|\|:|<=|=>|:)\s*["]{1,2}([^\s\{}?%[\]]+)["]{1,2}
    var passwdKey string = "zxcv0987bnml6543"

    // Exclude
    // (?i)(password|passwd|user_pwd)[a-z0-9_]*['\"]{0,2}\s*(=|>|:=|\||: |<=|=>|:)\s*((True|true|false|False|0|1)|['\"](True|true|false|False|0|1)['\"]|(['\"]{1,2}([^\s\{\}\?\%\[\]]+)['\"]{1,2}[\.]))
    password := "False"


    // SECRET: Generic Secret Key Credentials
    // (?i)['"]?\bsecret[a-z0-9_]*key[a-z0-9_]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(:)\s*([A-Za-z0-9/+=]{16,128})
    var secretKeyToken string = "abcd1234efgh5678ijkl9012"

    // SECRET: Generic Secret Key Credentials
    // (?i)['"]?\bsecret[a-z0-9_]*key[a-z0-9_]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(=|:=|=>|:)\s*['"]([a-zA-Z0-9_\-\.\$@*\^+!\[\]\(\)^=]{16,128})['"]
    var secretKeyQuoted string = "zxcv0987bnml6543.asdf@123"

    // Excluded by: (?i)['"]?\bsecret[a-z0-9_]*key[a-z0-9_]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(=|:=|=>|:)\s*['"]\$([a-zA-Z0-9_\-\.\$@*\^+!\[\]\(\)^=]{15,128})['"]
    var secretKeyExcluded string = "$zxcv0987bnml6543.asdf@123"


    // SECRET: Github Credentials
    // ['"]gh[pousr]_[0-9a-zA-Z]{32,40}['"]
    var github_token = "ghp_abcdefghijABCDEFGHIJabcdefghij123456"

    // SECRET: Github Credentials
    // (=|>|:=|\||: |<=|=>|:)\s*gh[pousr]_[0-9a-zA-Z]{32,40}
    var github_token_raw = ghs_abcdefghijABCDEFGHIJabcdefghij123456

    // SECRET: Github Credentials
    // ['"]?github_pat_[0-9a-zA-Z_\-]{60,}['"]?
    var github_pat = "github_pat_abcdefghijABCDEFGHIJ1234567890abcdefghijABCDEFGHIJ1234567890abcdefghijABCDEFGHIJ12"


    // SECRET: Gitlab Credentials
    // ['"]gl(pat|cict)-[A-Za-z0-9_-]{20,}['"]
    var gitlab_token = "glpat-abcdefghijABCDEFGHIJabcdefghij123456"

    // SECRET: Gitlab Credentials
    // (=|>|:=|\||: |<=|=>|:)\s*gl(pat|cict)-[A-Za-z0-9_-]{20,}
    var gitlab_token_raw = glcict-abcdefghijABCDEFGHIJabcdefghij123456


    // SECRET: Go Cardless Credentials
    // ['"]live_(?i)[a-z0-9\-_=]{40}['"]
    var gocardless_token = "live_abcdefghijABCDEFGHIJabcdefghij1234567890"

    // SECRET: Go Cardless Credentials
    // (=|>|:=|\||: |<=|=>|:)\s*live_(?i)[a-z0-9\-_=]{40}
    var gocardless_token_raw = live_abcdefghijABCDEFGHIJabcdefghij1234567890


    // SECRET: Grafana Credentials
    // (?i)grafana[a-z0-9_]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(=|>|:=|\|:|<=|=>|:)\s*['"]?eyJrIjoi(?i)[a-z0-9\-_=]{72,126}['"]?
    var grafanaToken string = "eyJrIjoiabcd1234efgh5678ijkl9012mnop3456qrst7890uvwxyz1234567890abcdef1234567890abcdefghij"

    // SECRET: Grafana Credentials
    // (?i)grafana[a-z0-9_]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(=|>|:=|\|:|<=|=>|:)\s*['"]?glsa_[a-zA-Z0-9_\-]{36,}['"]?
    var grafanaKey string = "glsa_abcd1234efgh5678ijkl9012mnop3456qrst7890"


    // SECRET: Hashicorp Credentials
    // ['"](?i)[a-z0-9]{10,24}.atlasv1.[a-z0-9-_=]{60,70}["']
    var hashicorp_secret_1 = "abcdefghij.atlasv1.1234567890ABCDEFGHIJabcdefghij1234567890ABCDEFGHIJabcdefghij"

    // SECRET: Hashicorp Credentials
    // (=|>|:=|\\|\\|:|<=|=>|:)\s*(?i)[a-z0-9]{10,24}\\.atlasv1\\.[a-z0-9\\-_=]{60,70}
    var hashicorp_secret_2 = abcdefghij.atlasv1.1234567890ABCDEFGHIJabcdefghij1234567890ABCDEFGHIJabcdefghij

    // SECRET: Hashicorp Credentials
    // ['"]tfp_[A-Za-z0-9]{20,40}["']
    var hashicorp_secret_3 = "tfp_abcdefghij1234567890ABCD"

    // SECRET: Hashicorp Credentials
    // (=|>|:=|\\|\\|:|<=|=>|:)\s*tfp_[A-Za-z0-9]{20,40}
    var hashicorp_secret_4 : tfp_abcdefghij1234567890ABCD


    // SECRET: Heroku Credentials
    // (?i)heroku[a-z0-9_\]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(=|>|:=|\|:|<=|=>|:)\s*['"]?([0-9A-F]{8}-[0-9A-F]{4}-[0-9A-F]{4}-[0-9A-F]{4}-[0-9A-F]{12})['"]?
    var herokuToken string = "12345678-abcd-1234-abcd-1234567890ab"

    // SECRET: Heroku Credentials
    // (?i)heroku[a-z0-9_\]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(=|>|:=|\|:|<=|=>|:)\s*['"]?(HRKU-[a-zA-Z0-9\_\-=+]{48,72})['"]?
    var herokuHrkToken string = "HRKU-abcd1234efgh5678ijkl9012mnop3456qrst7890uvwxyz1234567890abcd"

    // SECRET: Heroku Credentials
    // (?i)heroku[a-z0-9_\]*key[a-z0-9_\]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(=|>|:=|\|:|<=|=>|:)\s*['"]?([A-Za-z0-9/+=]{16,48})['"]?
    var herokuKey string = "abcd1234efgh5678ijkl9012"


    // SECRET: Htpasswd Credentials
    // ^[a-zA-Z0-9_-]+:\$apr1\$[A-Za-z0-9./]{8}\$[A-Za-z0-9./]{22}$
    var htpasswdValue string = "user1:$apr1$abcdefgh$1234567890123456789012"

    // SECRET: Htpasswd Credentials
    // (?i)[a-z0-9_]*htpasswd[a-z0-9_]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(=|>|:=|\|:|<=|=>|:)\s*['"]?([a-zA-Z0-9_-]+:\$apr1\$[A-Za-z0-9./]{8}\$[A-Za-z0-9./]{22})['"]?
    var userHtpasswd string = "admin:$apr1$ijklmnop$0987654321098765432109"


    // SECRET: Http Auth Header Credentials
    // (?i)Authorization[\\]?['"]?\s*(=|:=|:|=>)\s*[\\]?['"]?Basic\s+[A-Za-z0-9+/]{16,}={0,2}['"]?
    var authHeaderBasic string = "Authorization: Basic YWRtaW46c2VjdXJlcGFzczEyMw=="

    // SECRET: Http Auth Header Credentials
    // (?i)Authorization[\\]?['"]?\s*(=|:=|:|=>)\s*[\\]?['"]?Bearer\s+([A-Za-z0-9-._~+/]{20,})\s*['"]?
    var authHeaderBearer string = "Authorization: Bearer abcdefghijklmnopqrst1234567890"


    // SECRET: Hubspot Credentials
    // (?i)hubspot[a-z0-9_]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(=|>|:=|\|:|<=|=>|:)\s*['"]([a-z0-9]{8}-[a-z0-9]{4}-[a-z0-9]{4}-[a-z0-9]{4}-[a-z0-9]{12})['"]
    var hubspotToken string = "abcd1234-efgh-5678-ijkl-9012mnop3456"

    // SECRET: Hubspot Credentials
    // (?i)hubspot[a-z0-9_]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(=|>|:=|\|:|<=|=>|:)\s*['"]pat-[a-z0-9]{1,4}-[a-z0-9]{8}-[a-z0-9]{4}-[a-z0-9]{4}-[a-z0-9]{4}-[a-z0-9]{12}['"]
    var hubspotPat string = "pat-ab12-abcd1234-efgh-5678-ijkl-9012mnop3456"


    // SECRET: Hugging Face Token Credentials
    // ['"](hf_|api_org_|hf_app_)[a-zA-Z0-9]{30,64}['"]
    var huggingFaceToken string = "hf_abcd1234efgh5678ijkl9012mnop3456"

    // SECRET: Hugging Face Token Credentials
    // (=|>|:=|\|:|<=|=>|:)\s*(hf_|api_org_|hf_app_)[a-zA-Z0-9]{30,64}
    var huggingFaceKey string = api_org_1234567890abcdefeior1234567890ab


    // SECRET: Intercom Credentials
    // (?i)intercom[a-z0-9_ .,-]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(=|>|:=|\|:|<=|=>|:)\s*['"]([A-Za-z0-9=\-_]{48,72})['"]
    var intercomToken string = "abcd1234efgh5678ijkl9012mnop3456qrst7890uvwxyz1234567890abcdefghij"

    // SECRET: Intercom Credentials
    // (?i)intercom[a-z0-9_ .,-]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(=|>|:=|\|:|<=|=>|:)\s*['"]([a-h0-9]{8}-[a-h0-9]{4}-[a-h0-9]{4}-[a-h0-9]{4}-[a-h0-9]{12})['"]
    var intercomUuid string = "abcd1234-efgh-5678-abcd-9012efgh3456"


    // SECRET: Ionic Credentials
    // (?i)ionic[a-z0-9_ .,-]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(=|>|:=|\|:|<=|=>|:)\s*['"](ion_[A-Za-z0-9]{32,45})['"]
    var ionicToken string = "ion_abcd1234efgh5678ijkl9012mnop3456qrst"


    // SECRET: Jekyll Token Credentials
    // (?i)[a-z0-9_]*jekyll[a-z0-9_]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(=|>|:=|\|:|<=|=>|:)\s*[a-zA-Z0-9]{24,}
    var jekyllToken string = abcd1234efgh5678ijkl9012mnop3456

    // SECRET: Jekyll Token Credentials
    // (?i)[a-z0-9_]*jekyll[a-z0-9_]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(=|>|:=|\|:|<=|=>|:)\s*['"][a-zA-Z0-9]{24,}['"]
    var jekyllKey string = "1234567890abcdef1234567890abcdef"


    // SECRET: Jwt Token Credentials
    // \b(?i)jwt[a-z0-9_]*['"]?\s*(=|:|=>|:=)\s*['"]([A-Za-z0-9-_]+\.[A-Za-z0-9-_]+\.[A-Za-z0-9-_]+)['"]
    var jwtToken string = "eyJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJ1c2VyIn0.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"

    // SECRET: Jwt Token Credentials
    // eyJ[a-zA-Z0-9_\-]+\.(eyJ[A-Za-z0-9-_]+)\.[A-Za-z0-9-_]+
    var jwtAuth string = "eyJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJ1c2VyIn0.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"

    // SECRET: Jwt Token Credentials
    // eyJ[a-zA-Z0-9_\-]+\.[A-Za-z0-9-_]+\.[A-Za-z0-9-_]+\.[A-Za-z0-9-_]+\.[A-Za-z0-9-_]+
    var jwtExtended string = "eyJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJ1c2VyIn0.abcd1234.efgh5678.ijkl9012"


    // SECRET: Launchdarkly Token Credentials
    // ['"]api-[a-z0-9]{8}-[a-z0-9]{4}-[a-z0-9]{4}-[a-z0-9]{4}-[a-z0-9]{12}['"]
    var launchDarklyToken string = "api-abcd1234-efgh-5678-ijkl-9012mnop3456"

    // SECRET: Launchdarkly Token Credentials
    // ['"]api-[a-f0-9]{32}['"]
    var launchDarklyKey string = "api-abcdef0123456789abcdef0123456789"

    // SECRET: Launchdarkly Token Credentials
    // (=|>|:=|\|:|<=|=>|:)\s*api-[a-z0-9]{8}-[a-z0-9]{4}-[a-z0-9]{4}-[a-z0-9]{4}-[a-z0-9]{12}
    var launchDarklyUuid string = api-abcd1234-efgh-5678-ijkl-9012mnop3456

    // SECRET: Launchdarkly Token Credentials
    // (=|>|:=|\|:|<=|=>|:)\s*api-[a-f0-9]{32}
    var launchDarklyHex string = api-1234567890abcdef1234567890abcdef


    // SECRET: Linear Credentials
    // ['"]lin_api_[A-Za-z0-9]{32,44}['"]
    var linearApiToken string = "lin_api_abcd1234efgh5678ijkl9012mnop3456"

    // SECRET: Linear Credentials
    // (=|>|:=|\|:|<=|=>|:)\s*lin_api_[A-Za-z0-9]{32,44}
    var linearApiKey string = lin_api_1234567890abcdef1234567890abcdef

    // SECRET: Linear Credentials
    // (?i)linear[a-z0-9_ .,-]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(=|>|:=|\|:|<=|=>|:)\s*['"]([a-f0-9]{28,44})['"]
    var linearToken string = "abcdef0123456789abcdef0123456789"


    // SECRET: Linkedin Credentials
    // (?i)linkedin[a-z0-9_]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(=|>|:=|\|:|<=|=>|:)\s*['"]((?i)[a-z0-9_\.=+!-]{14,64})['"]
    var linkedinToken string = "abcd1234efgh5678_ijkl9012.mnop=qr+st"


    // SECRET: Lob Credentials
    // (?i)lob[a-z0-9_ .,-]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(=|>|:=|\|:|<=|=>|:)\s*['"]((live|test)_[a-f0-9]{28,38})['"]
    var lobToken string = "live_abcdef0123456789abcdef0123456789"

    // SECRET: Lob Credentials
    // (?i)lob[a-z0-9_ .,-]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(=|>|:=|\|:|<=|=>|:)\s*['"]((live|test)_pub_[a-f0-9]{28,35})['"]
    var lobPubToken string = "live_pub_1234567890abcdef1234567890ab"


    // SECRET: Mailchimp Token Credentials
    // (?i)[a-z0-9_]*mailchimp[a-z0-9_]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(=|>|:=|\|:|<=|=>|:)\s*[a-zA-Z0-9]{24,42}-us[0-9]{1,3}
    var mailchimpToken string = abcd1234efgh5678ijkl9012mnop3456-us12

    // SECRET: Mailchimp Token Credentials
    // (?i)[a-z0-9_]*mailchimp[a-z0-9_]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(=|>|:=|\|:|<=|=>|:)\s*['"][a-zA-Z0-9]{24,42}-us[0-9]{1,3}['"]
    var mailchimpKey string = "1234567890abcdef1234567890abcdef-us12"


    // SECRET: Mailgun Credentials
    // (?i)mailgun[a-z0-9_]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(=|>|:=|\|:|<=|=>|:)\s*['"]((pub)?key-[a-f0-9]{32})['"]
    var mailgunToken string = "key-abcdef0123456789abcdef0123456789"

    // SECRET: Mailgun Credentials
    // (?i)mailgun[a-z0-9_]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(=|>|:=|\|:|<=|=>|:)\s*['"]([a-h0-9]{32}-[a-h0-9]{8}-[a-h0-9]{8})['"]
    var mailgunKey string = "abcdef0123456789abcdef0123456789-abcdefgh-abcdefgh"


    // SECRET: Mapbox Credentials
    // (?i)(pk\.[a-z0-9\-_]{54,100}\.[a-z0-9\-_]{16,32})
    var mapboxToken string = "pk.abcd1234efgh5678ijkl9012mnop3456qrst7890uvwxyz1234567890ab.abcdefghijklmnop"


    // SECRET: Messagebird Credentials
    // (?i)messagebird[a-z0-9_]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(=|>|:=|\|:|<=|=>|:)\s*['"]((?i)[a-z0-9]{20,48})['"]
    var messagebirdToken string = "abcd1234efgh5678ijkl9012mnop3456"

    // SECRET: Messagebird Credentials
    // (?i)messagebird[a-z0-9_]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(=|>|:=|\|:|<=|=>|:)\s*['"]((?i)[a-h0-9]{8}-[a-h0-9]{4}-[a-h0-9]{4}-[a-h0-9]{4}-[a-h0-9]{12})['"]
    var messagebirdUuid string = "abcd1234-efgh-5678-abcd-9012efgh3456"


    // SECRET: Mongodb Connection String Credentials
    // \bmongodb(?:\+srv)?:\/\/(?!<)[^\s:@\/]+:[^\s:@\/]+@[^\s]+
    var mongoDbConnection string = "mongodb://admin:securepass123@cluster0.mongodb.net"


    // SECRET: Mysql Connection String Credentials
    // (?i)`?(mysql|jdbc:mysql)://((?!{.*})[^\s:@/]+)(:((?!{.*})[^\s:@/]+))?@[^\\s:/]+(:[^\\s/]+)?/((?!{.*})[^\\s?]+)`?
    var mysqlConnection string = "mysql://admin:securepass123@localhost:3306/mydb"

    // SECRET: Mysql Connection String Credentials
    // (?i)`?(mysql|jdbc:mysql)://[^\\s:/]+(:[^\\s/]+)?/((?!{.*})[^\\s?]+)\?.*?((user=((?!{.*})[^&\s]+))|(password=((?!{.*})[^&\s]+)))
    var mysqlConnectionParams string = "mysql://localhost:3306/mydb?user=admin&password=securepass123"


    // SECRET: New Relic Credentials
    // NRAK-[A-Z0-9]{16,30}\b
    var newRelicApiKey string = "NRAK-XYZ78901PQR23456STUV"

    // SECRET: New Relic Credentials
    // NRJS-[a-z0-9]{16,30}\b
    var newRelicJsKey string = "NRJS-xyzw7890pqrs1234tuv"

    // SECRET: New Relic Credentials
    // NRBR-[a-z0-9]{16,30}\b
    var newRelicBrowserKey string = "NRBR-9012345678xyzab12345"

    // SECRET: New Relic Credentials
    // (?i)newrelic[a-z0-9_]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(=|>|:=|\|:|<=|=>|:)\s*['"]((?i)[a-z0-9\-_]{44,82})['"]
    var newRelicToken string = "xyzw7890-pqrs1234-tuv5678-abcd9012-efgh3456-jklm"


    // SECRET: Npm Access Credentials
    // ['"]npm_(?i)[a-z0-9]{30,42}['"]
    var npmToken string = "npm_abcdef0123456789ghijklmnop1234567890"


    // SECRET: Openai Token Credentials
    // ['"]sk-[a-zA-Z0-9]{25,50}['"]
    var openAiToken string = "sk-abcd1234EFGH5678ijkl9012MNOP3456"

    // SECRET: Openai Token Credentials
    // (=|>|:=|\|:|<=|=>|:)\s*sk-[a-zA-Z0-9]{25,50}
    var openAiKey string = sk-12345678abcdEFGH9012ijklMNOP3456


    // SECRET: P12 Credentials
    // \.p12\s*(=|:=|:|=>|)\s*['"]([a-zA-Z0-9\-_./]{24,})['"]
    var p12Cert string = "cert.p12=abcd1234-efgh5678_ijkl9012.mnop3456"

    // SECRET: P12 Credentials
    // \.p12\s*(=|:=|:|=>|)\s*([a-zA-Z0-9\-_./]{24,})
    var p12Path string = cert.p12=12345678-abcd9012_efgh3456.ijkl7890


    // SECRET: Paypal Credentials
    // (?i)paypal[a-z0-9_ .,-]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(=|>|:=|\|:|<=|=>|:)\s*['"]([a-zA-Z0-9\-_]{24,84})['"]
    var paypalToken string = "abcd1234EFGH5678_ijkl9012-MNOP3456"

    // SECRET: Paypal Credentials
    // ['"]A21A[A-Za-z0-9-_]{92}['"]
    var paypalApiKey string = "A21A90123456XYZW7890abcd1234EFGH5678ijkl9012MNOPqrst3456UVWXyzab7890CDEF1234ghij5678klmn90122222"

    // SECRET: Paypal Credentials
    // (=|>|:=|\|:|<=|=>|:)\s*A21A[A-Za-z0-9-_]{92}
    var paypalApiSecret string = A21A90123456XYZW7890abcd1234EFGH5678ijkl9012MNOPqrst3456UVWXyzab7890CDEF1234ghij5678klmn90122222


    // SECRET: Planetscale Credentials
    // ['"]pscale_pw_(?i)[a-z0-9\-_\.]{36,52}['"]
    var planetScaleKey string = "pscale_pw_abcd1234-efgh5678_ijkl9012.mnop3456qrst"

    // SECRET: Planetscale Credentials
    // ['"]pscale_tkn_(?i)[a-z0-9\-_\.]{36,52}['"]
    var planetScaleToken string = "pscale_tkn_12345678-abcd9012_efgh3456.ijkl7890mnop"


    // SECRET: Postgresql Connection String Credentials
    // (?i)`?(postgres|jdbc:postgresql|postgresql)\:\/\/[^\s\:@\/]+(\:[^{}\s\:@\/]+)@[^{}\s\:\/]+(\:[^{}\s\/]+)?\/[^{}\s?]+`?
    var postgresConnection string = "postgres://admin:secure123@localhost:5432/mydb"

    // SECRET: Postgresql Connection String Credentials
    // (?i)`?(postgres|jdbc:postgresql|postgresql)\:\/\/[^\s\:\/]+(\:[^\s\/]+)?\/((?!{.*})[^\s?]+)\?.*?((user=((?!{.*})[^&\s]+))|(password=((?!{.*})[^&\s]+)))
    var postgresConnectionParams string = "postgres://localhost:5432/mydb?user=admin&key=secure123"


    // SECRET: Postman Api Credentials
    // PMAK-(?i)[a-f0-9]{24}-[a-f0-9]{34}\b
    var postman_api_key = "PMAK-abcdefabcd12345678901234-1234567890abcdefabcd1234567890abcd"


    // SECRET: Pulumni Api Credentials
    // pul-[a-f0-9]{40}\\b
    var pulumni_api_key = "pul-abcdeabcde1234567890abcdeabcde1234567890"


    // SECRET: Putty Key Credentials
    // (?si)PuTTY-User-Key-File-\\d+: \\S+\\s+((?![\"]).)*?Private-Lines: [0-9]+(?:\\s+[A-Za-z0-9+/=]+)+(?:\\s+Private-MAC: [a-f0-9]+)?
    var putty_key = "PuTTY-User-Key-File-3: ssh-rsa\nEncryption: none\nComment: user@host\nPrivate-Lines: 2\nabcdefghij1234567890ABCDEFGHIJ1234567890==\nabcdefghij1234567890ABCDEFGHIJ1234567890==\nPrivate-MAC: 1234567890abcdef1234567890abcdef1234567890"

    // SECRET: Putty Key Credentials
    // ['\""](?si)PuTTY-User-Key-File-\\d+:[^'\""]*Private-Lines: [0-9]+[^'\""]*Private-MAC: [a-f0-9]+['\""]
    var putty_key = "PuTTY-User-Key-File-3: ssh-rsa\nEncryption: none\nComment: user@host\nPrivate-Lines: 2\nabcdefghij1234567890ABCDEFGHIJ1234567890==\nabcdefghij1234567890ABCDEFGHIJ1234567890==\nPrivate-MAC: 1234567890abcdef1234567890abcdef1234567890"


    // SECRET: Pypi Credentials
    // ['\""]pypi-AgEIcHlwaS5vcmc[A-Za-z0-9_-]{24,}['\""]
    var pypi_api_key = "pypi-AgEIcHlwaS5vcmcabcdefghij1234567890-_34"

    // SECRET: Pypi Credentials
    // (=|>|:=|\|\||:|<|<=|=>|:)\\s*pypi-AgEIcHlwaS5vcmc[A-Za-z0-9_-]{24,}
    var pypi_api_key = pypi-AgEIcHlwaS5vcmcabcdefghij1234567890ABCD-_


    // SECRET: Rubygems Api Credentials
    // rubygems_[a-f0-9]{40,52}\\b
    var rubygems_api_key = "rubygems_abcdeabcde1234567890abcdeabcde1234567890"


    // SECRET: Sendgrid Api Credentials
    // SG\.(?i)[a-z0-9_\\-\\.]{54,78}\\b
    var sendgrid_api_key = "SG.abcdefghij1234567890abcdefghij1234567_-.abcdefghij1234"


    // SECRET: Sendinblue Api Credentials
    // xkeysib-[a-f0-9]{64}-(?i)[a-z0-9]{16}\\b
    var sendinblue_api_key = "xkeysib-abcdeabcde1234567890abcdeabcde1234567890abcdeabcde12345678901234-1234567890abcdef"


    // SECRET: Shippo Api Credentials
    // shippo_(live|test)_[a-f0-9]{36,48}\b
    var shippoApiKey string = "shippo_test_ffeeddccbb9876543210ffeeddccbb9876543210"


    // SECRET: Shopify Credentials
    // ['\""]shp(ss|at|ca|pa)_[a-zA-Z0-9]{32}['\""]
    var shopify_api_key_1 = "shpat_abcdefghij1234567890ABCDEFGHIJ12"

    // SECRET: Shopify Credentials
    // (=|>|:=|\\|\\|:|<=|=>|:)\\s*shp(ss|at|ca|pa)_[a-zA-Z0-9]{32}
    var shopify_api_key_2 = shpss_abcdefghij1234567890ABCDEFGHIJ12


    // SECRET: Slack Credentials
    // https://hooks.slack.com/services/[A-Za-z0-9+/]{40,}
    var slack_api_key_1 = "https://hooks.slack.com/services/abcdefghij1234567890ABCDEFGHIJ1234567890+/"

    // SECRET: Slack Credentials
    // ['\""]xox[baprs]-[0-9a-zA-Z\\-]{10,120}['\""]
    var slack_api_key_2 = "xoxb-abcdefghij1234567890ABCDEFGHIJ1234567890"

    // SECRET: Slack Credentials
    // (=|>|:=|\\|\\|:|<=|=>|:)\\s*xox[baprs]-[0-9a-zA-Z\\-]{10,120}
    var slack_api_key_3 = xoxa-abcdefghij1234567890ABCDEFGHIJ1234567890-


    // SECRET: Stripe Credentials
    // (?i)stripe[a-z0-9_]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(=|>|:=|\|:|<=|=>|:)\s*['"](?i)(sk|pk)_(live|test)_[a-z0-9]{10,120}['"]
    var stripeApiKey1 string = "sk_live_aabbccdeef0123456789"

    // SECRET: Stripe Credentials
    // (?i)stripe[a-z0-9_]*\s*(=|>|:=|\|:|<=|=>|:)\s*(?i)(sk|pk)_(live|test)_[a-z0-9]{10,120}
    var stripeApiKey2 = pk_test_0123456789aabbccdeef

    // SECRET: Stripe Credentials
    // ['"]sk_(live|test)_[A-Za-z0-9]{10,32}['"]
    var stripeApiKey3 string = "sk_live_AABBCCDDEF0123456789"

    // SECRET: Stripe Credentials
    // (=|>|:=|\|:|<=|=>|:)\s*sk_(live|test)_[A-Za-z0-9]{10,32}
    var stripeApiKey4 string = sk_test_AABBCCDDEF0123456789


    // SECRET: Twilio Credentials
    // ['"]SK[0-9a-fA-F]{32}['"]
    var twilioApiKey1 string = "SKaabbccdeef0123456789abcdef01234567"

    // SECRET: Twilio Credentials
    // (=|>|:=|\|:|<=|=>|:)\s*SK[0-9a-fA-F]{32}
    var twilioApiKey2 string = SKaabbccdeef0123456789abcdef01234567

    // SECRET: Twilio Credentials
    // (?i)twilio[a-z0-9_]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(=|>|:=|\|:|<=|=>|:)\s*['"]([A-Za-z0-9+/]{24,50})['"]
    var twilioApiKey3 string = "aabbccdeef0123456789abcd"


    // SECRET: Twitch Credentials
    // (?i)twitch[a-z0-9_]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(=|>|:=|\|:|<=|=>|:)\s*['"](?i)[a-z0-9]{24,36}['"]
    var twitchApiKey string = "aabbccdeef0123456789aabbccdeef"


    // SECRET: Twitter Credentials
    // (?i)twitter[a-z0-9_]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(=|>|:=|\|:|<=|=>|:)\s*['"]([A-Za-z0-9+/\-_+]{24,64})['"]
    var twitterApiKey string = "aabbccdeef0123456789abcdABCD12"


    // SECRET: Typeform Api Credentials
    // (?i)typeform[a-z0-9_]*(?:[\:\s]+)?(?:string|String|str)?['"]?\s*(=|>|:=|\|:|<=|=>|:)\s*['"](tfp_[a-z0-9\-_\.=]{54,64})['"]
    var typeformApiKey string = "tfp_aabbccdeef0123456789_.-=aabbccdeef0123456789_.-=aabbcc"


    // SECRET: Wordpress Auth Credentials
    // define(s*['"]*(DB_PASSWORD|MYSQL_PASSWORD|DB_PASS|DATABASE_PASSWORD)['"]*s*,s*['"][^'"]+['"]s*)
    var wordpress_api_key_1 = "define('DB_PASSWORD', 'abcdefghij1234567890ABCDEFGHIJ12!@#$')"

    // SECRET: Wordpress Auth Credentials
    // define(s*['"]*(AUTH_KEY|SECURE_AUTH_KEY|LOGGED_IN_KEY|NONCE_KEY|AUTH_SALT|SECURE_AUTH_SALT|LOGGED_IN_SALT|NONCE_SALT)['"]*s*,s*['"][^'"]+['"]s*)
    var wordpress_api_key_2 = "define('AUTH_KEY', 'abcdefghij1234567890ABCDEFGHIJ12!@#$')"


    // SECRET: Zoom Credentials
    // (?i)zoom[a-z0-9_ .\-,]*(?:[\:\s]+)?(?:string|String|str)?['\"]?\s*(=|>|:=|\|\|:|<=|=>|:)\s*['\"]([a-zA-Z0-9\-_]{24,60})['\"]
    var zoom_token string = "AbcdEFGH1234ijkl-5678MNOP_qrstUVWX90yz12"
}
<!doctype html>
<html lang='{{.Lang}}'>

<head>
    <title>{{.Title}}</title>
    <meta name='viewport' content='width=device-width, initial-scale=1, shrink-to-fit=no' />
    <meta charset='utf-8' />
    <base target='_blank' />

    <link rel='icon' href='/favicon.ico' />

    <link rel='stylesheet' href='//cdnjs.cloudflare.com/ajax/libs/twitter-bootstrap/4.1.3/css/bootstrap.min.css' />

    <style>
        .center {
            text-align: center;
        }

        .thumb {
            width: 100%;
        }

        .room {
            background-color: whitesmoke;
            margin: 1em;
        }

        .column {
            -webkit-column-width: 200px;
            -moz-column-width: 200px;
            column-width: 200px;
            -webkit-column-count: 3;
            -moz-column-count: 3;
            column-count: 3;
            -webkit-column-gap: 30px;
            -moz-column-gap: 30px;
            column-gap: 30px;
            -webkit-column-rule: 1px solid #ccc;
            -moz-column-rule: 1px solid #ccc;
            column-rule: 1px solid #ccc;
        }

        .video_wrapper {
            margin: 0 auto;
            position: relative;
        }

        .videoPoster {
            position: absolute;
            top: 0px;
            left: 0px;
            width: 100%;
            height: 100%;
            background-color: transparent;
            background-size: 100%;
            border: 1px solid grey;
            overflow: hidden;
            opacity: 1;
            -webkit-transition: opacity 800ms, height 0s;
            -moz-transition: opacity 800ms, height 0s;
            transition: opacity 800ms, height 0s;
            -webkit-transition-delay: 0s, 0s;
            -moz-transition-delay: 0s, 0s;
            transition-delay: 0s, 0s;
        }

        .videoPoster:hover {
            cursor: pointer;
        }

        .videoPoster:before {
            content: '';
            position: absolute;
            top: 50%;
            left: 50%;
            width: 80px;
            height: 80px;
            margin: -40px 0 0 -40px;
            border: 5px solid #fff;
            border-radius: 100%;
            -webkit-transition: border-color 300ms;
            -moz-transition: border-color 300ms;
            transition: border-color 300ms;
        }

        .videoPoster:after {
            content: '';
            position: absolute;
            top: 48%;
            left: 49%;
            width: 0;
            height: 0;
            margin: -20px 0 0 -10px;
            border-left: 40px solid #fff;
            border-top: 25px solid transparent;
            border-bottom: 25px solid transparent;
            -webkit-transition: border-color 300ms;
            -moz-transition: border-color 300ms;
            transition: border-color 300ms;
        }

        .videoPoster:hover:before,
        .videoPoster:focus:before {
            border-color: #f00;
        }

        .videoPoster:hover:after,
        .videoPoster:focus:after {
            border-left-color: #f00;
        }

        .videoWrapperActive .videoPoster {
            opacity: 0;
            height: 0;
            -webkit-transition-delay: 0s, 800ms;
            -moz-transition-delay: 0s, 800ms;
            transition-delay: 0s, 800ms;
        }
    </style>
</head>

<body style='background-color: {{.Color}};'>

    <main role='main'>

        <div class='jumbotron'>
            <div class='container'>
                <div class='row'>
                    <div class='col-md-4'>
                        <h1 class='center'>{{.Room.DisplayName}}</h1>
                        <div class="video_wrapper video_wrapper_full">
                            <a href='{{out .Room.UserName}}'>
                                <img src='{{image .Room.ProfileImages.ThumbnailImageBig}}' class='thumb' />
                                <button class="videoPoster"></button>
                            </a>
                        </div>
                        <h2 class='center'>{{.Room.DisplayAge}}</h2>
                        <h3 class='center'>{{.Room.HomeTown}}</h3>
                        <div class='clearfix'></div>
                    </div>
                    <div class='col-md-8'>
                        <div class='column'>
                            <p>
                                <strong>Дата регистрации</strong>: {{.Room.SignupDate}}
                            </p>
                            <p>
                                <strong>Последний логин </strong>: {{.Room.LastUpdateDate}}
                            </p>
                            <p>
                                <strong>Время онлайн</strong>: {{timeonline .Room.OnlineTime}}
                            </p>
                            <p>
                                <strong>Подписчиков</strong>: {{.Room.MembersCount}}
                            </p>
                            <p>
                                <strong>Лицо</strong>: {{.Room.Ethnicity}}
                            </p>
                            <p>
                                <strong>Язык</strong>: {{.Room.PrimaryLanguage}}, {{.Room.SecondaryLanguage}}
                            </p>
                            <p>
                                <strong>Нравится</strong>: {{.Room.TurnsOn}}
                            </p>
                            <p>
                                <strong>Не нравится</strong>: {{.Room.TurnsOff}}
                            </p>
                            <p>
                                <strong>Рост</strong>: {{height .Room.Height}}
                            </p>
                            <p>
                                <strong>Вес</strong>: {{weight .Room.TurnsOff}}
                            </p>
                            <p>
                                <strong>Волосы</strong>: {{.Room.HairColor}}
                            </p>
                            <p>
                                <strong>Глаза</strong>: {{.Room.EyeColor}}
                            </p>
                            <p>
                                <strong>Грудь</strong>: {{size .Room.BustPenisSize}}
                            </p>
                            <p>
                                <strong>Интимная стрижка</strong>: {{.Room.PubicHair}}
                            </p>
                        </div>
                    </div>
                </div>
            </div>
        </div>

        <div class='container'>
            <div class='row'>
                {{range .Rooms}}
                <div class='col-md-3' class='room'>
                    <a href='{{link .UserName}}'>
                        <img src='{{image .ProfileImages.ThumbnailImageMedium}}' class='thumb' />
                    </a>
                    <h4 class='center'>{{.DisplayName}}</h4>
                </div>
                {{end}}
            </div>
        </div>
    </main>

    <footer class='container'>
        <p class='center'>&copy; {{.Domain}} 2018</p>
    </footer>

    <script src='//code.jquery.com/jquery-3.2.1.min.js'></script>
    <script src='//cdnjs.cloudflare.com/ajax/libs/twitter-bootstrap/4.1.3/js/bootstrap.min.js'></script>
</body>

</html>
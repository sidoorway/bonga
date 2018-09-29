<!doctype html>
<html lang='{{.Lang}}'>

<head>
    <title>{{.Title}}</title>
    <meta name='viewport' content='width=device-width, initial-scale=1, shrink-to-fit=no' />
    <meta charset='utf-8' />
    <base target='_blank' />

    <link rel='icon' href='/favicon.ico' />

    <!-- Bootstrap core CSS -->
    <link rel='stylesheet' href='//cdnjs.cloudflare.com/ajax/libs/twitter-bootstrap/4.1.3/css/bootstrap.min.css' />

    <!-- Custom styles for this template -->
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
    </style>
</head>

<body style='background-color: {{.Color}};'>

    <main role='main'>
        {{range .Genders}}
        <div class='container'>
            <hr/>
            <h2 class='center'>{{.Title}}</h2>
            <hr />
            <div class='row'>
                {{range .Rooms}}
                <div class='col-md-3' class='room'>
                    <a href='{{link .UserName}}'>
                        <img src='{{image .ProfileImages.ThumbnailImageMedium}}' class='thumb' />
                    </a>
                    <h4 class='center'><strong>{{.DisplayName}}</strong></h4>
                </div>
                {{end}}
            </div>
        </div>
        {{end}}
    </main>

    <footer class='container'>
        <p class='center'>&copy; {{.Domain}} 2018</p>
    </footer>

    <!-- Bootstrap core JavaScript
    ================================================== -->
    <!-- Placed at the end of the document so the pages load faster -->
    <script src='//code.jquery.com/jquery-3.2.1.slim.min.js'></script>
    <script src='//cdnjs.cloudflare.com/ajax/libs/popper.js/1.14.4/esm/popper.min.js'></script>
    <script src='//cdnjs.cloudflare.com/ajax/libs/twitter-bootstrap/4.1.3/js/bootstrap.min.js'></script>
</body>

</html>
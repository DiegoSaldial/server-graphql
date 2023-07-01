<!DOCTYPE html>
<html lang="es">
<head>
    <title>{{ .Title }}</title>
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">  

    <link href="https://fonts.googleapis.com/css?family=Roboto:100,300,400,500,700,900|Material+Icons" rel="stylesheet" type="text/css">
    <link href="https://cdn.jsdelivr.net/npm/quasar@2.12.0/dist/quasar.prod.css" rel="stylesheet" type="text/css"> 
    <script src="https://cdn.jsdelivr.net/npm/vue@3/dist/vue.global.prod.js"></script> 
    <script src="https://unpkg.com/@lottiefiles/lottie-player@latest/dist/lottie-player.js"></script>
    <script type="importmap">
        {
            "imports": {
                "vue": "https://unpkg.com/vue@3/dist/vue.esm-browser.js"
            }
        }
    </script>
    
    {{ block "css" . }}{{ end }}
</head>
<body>

    <div id="app">

        <div class="">
            <q-layout view="hHh Lpr lff" container style="height: 100vh" class="shadow-0">
            <q-header class="bg-primary">
                <q-toolbar>
                <q-toolbar-title>
                    <q-img src="https://cdn.quasar.dev/logo-v2/svg/logo-dark.svg" alt="logo" :ratio="1" style="
                        height: 30px;
                        max-width: 30px;
                        margin-right: 0.5em;
                        border-radius: 0%;
                    " />
                    <span style="margin-right: 0.7em" @click="$router.push('/')">
                    Webinar
                    </span>
                </q-toolbar-title>
                <q-toggle v-model="$q.dark.isActive" color="white" />
                
                </q-toolbar>
            </q-header>

            <q-page-container>
                <div class="container">
                    {{ block "content" . }}{{ end }}
                </div>
            </q-page-container>
            
            </q-layout>
        </div>

    </div>

    <script src="https://cdn.jsdelivr.net/npm/quasar@2.12.0/dist/quasar.umd.prod.js"></script>
    <script src="https://cdn.jsdelivr.net/npm/quasar@2.12.0/dist/lang/es.umd.prod.js"></script>
    

    {{ block "js" . }} {{ end }}
</body>
</html>
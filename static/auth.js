const { createApp, ref } = Vue 

const app = createApp({
  setup() {
    const loadingpdf = ref(false)
    const loadingw = ref(false)
    const camposvacios = ref(false) 
    const documento = ref("")

    function descargarPdf(){
      console.log("descargarPdf",Quasar)
    } 

    function goLogin(){
        console.log("login...");
        window.location.href = "/login";
    }

    return { 
        loadingpdf,
        loadingw,
        camposvacios, 
        documento,
        descargarPdf,
        goLogin,
    }
  }
});
app.use(Quasar)
Quasar.lang.set(Quasar.lang.es)
app.mount('#app')


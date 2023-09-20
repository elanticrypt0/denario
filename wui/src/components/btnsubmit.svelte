<script lang="ts">

    import apiURL from '../api/api';

    export let btnText="Enviar";
    export let btnTextOnSubmitting="Enviando..."
    export let goback:boolean=false;

    export let url="";
    export let object={};
    export let method="POST";

    let is_disabled=false;
    let btnTextShow=btnText;

    async function submitToApi(){

        is_disabled=true;
        btnTextShow=btnTextOnSubmitting;
        // object=object();
        const data=object();
        console.log(data);
        console.log('enviando a '+apiURL+url);
        console.log(data);

        if(data!=false){
            const response= await fetch(apiURL+url,{
                method:method,
                body:data,
                headers:{
                    'Content-Type':'application/json'
                }
            });
        }

        is_disabled=false;
        btnTextShow=btnText;

        if(goback){
            // const go2to=window.location.protocol+"//"+window.location.hostname+":"+window.location.port+"/"+url;
            // window.location;
            history.back();
        }
    }


</script>

<button type="button" id="submitButton" class="py-2 px-4  bg-green-700 hover:bg-green-400 focus:ring-indigo-500 focus:ring-offset-indigo-200 text-white w-full transition ease-in duration-200 text-center text-base font-semibold shadow-md focus:outline-none focus:ring-2 focus:ring-offset-2  rounded-lg " on:click={submitToApi} disabled='{is_disabled}'>
    {btnTextShow}
</button>
<script lang="ts">

    import { onMount } from 'svelte';
    import apiURL from '../../api/api';
    
    import Btnsubmit from '../btnsubmit.svelte';
    

    export let is_update=false;
    
    let name="";
    let submitMethod="POST";
    let endpoint="categories";

    function getNewCategorie():string{
        return JSON.stringify({
            name,
        });
    }
    
    function resetForm(){
        name="";        
    }

    onMount(async () =>{       
        if(is_update){
            submitMethod="PUT";
            // obtengo el ID del record del location hash

            const hash=window.location.hash;
            const category_id=hash.substring(1);
            
            endpoint+="/"+category_id;
            
            // lo cargo en el formulario
            const response=await fetch(apiURL+endpoint);
            const category_data=await response.json();            

            name=category_data.name;

        }
    });

</script>


<div class="flex-l mx-52 bg-white px-10 pt-5 pb-5 rounded-md">
    <div class="">
        <div class="relative mt-1 rounded-md shadow-sm">
            <div class="absolute inset-y-0 left-0 flex items-center pl-3 pointer-events-none">
                <span class="text-gray-500 sm:text-sm">
                    >
                </span>
            </div>
            <input type="text" name="name" id="name" class="block w-full px-4 py-2 pr-12 border-t border-b border-l border-gray-300 rounded-md focus:ring-indigo-500 focus:border-indigo-500 pl-7 sm:text-sm" placeholder="Nombre de la categoría..." bind:value={name} autocomplete="off"/>
            <div class="absolute inset-y-0 right-0 flex items-center">
                <label for="name" class="sr-only">
                    Nombre
                </label>
            </div>
        </div>
    </div>

    <div class="mt-10">
        <Btnsubmit btnText="Guardar" btnTextOnSubmitting="Guardando..." url={endpoint} method={submitMethod} object={getNewCategorie} goback={true}></Btnsubmit>
    </div>
</div>
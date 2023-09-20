<script lang="ts">

    import { onMount } from 'svelte';
    import apiURL from '../../api/api';
    
    import { ApiCategories } from '../../api/categories';
    import { DateHelper } from '../../helpers/datehelper';
    import { ApiAccounts } from '../../api/accounts';
    import Btnsubmit from '../btnsubmit.svelte';
    

    export let is_update=false;

    let now_today=DateHelper.getFullDate();

    let selected_amount_io:any="out";
    let amount_io = [
        { value: 'out', name: 'Egreso', default:true },
        { value: 'in', name: 'Ingreso' },
    ];
    
    let selected_category:any;
    let categories:any[] | any=[];
    
    let selected_account:any;
    let accounts = [
        { value: 'out', name: 'Egreso', default:true },
        { value: 'in', name: 'Ingreso' }
    ];

    let name="";
    let amount:string;
    let comment="-";
    let is_mutable="true";
    let submitMethod="POST";
    let endpoint="records";


    async function getCategories(){
        const response=await fetch(apiURL+"categories");
        const data=await response.json();
        categories=data;
    }
    
    async function getAccounts(){
        const response=await fetch(apiURL+"accounts");
        const data=await response.json();
        accounts=data;
    }


    function getNewRecord():string{
        return JSON.stringify({
            name,
            amount:Number(amount),
            amount_io:selected_amount_io,
            comment,
            record_date:now_today,
            category_id:selected_category,
            is_mutable:Boolean(is_mutable),
            account_id:selected_account
        });
    }
    

    function resetForm(){
        name="";
        amount="0.00";
        comment="-";
        now_today=DateHelper.getFullDate();
        selected_category=23;
        is_mutable="true";
        selected_account=3;
    }

    onMount(async () =>{
        await getCategories();
        await getAccounts();
        
        if(is_update){
            submitMethod="PUT";
            // obtengo el ID del record del location hash

            const hash=window.location.hash;
            const record_id=hash.substring(1);

            endpoint+="/"+record_id;

            // lo cargo en el formulario

            const response=await fetch(apiURL+"records/"+record_id);
            const record_data=await response.json();
            
            amount=record_data.amount;
            selected_amount_io=record_data.amount_io;
            name=record_data.name;
            comment=record_data.comment;
            now_today=record_data.record_date;
            selected_category=record_data.category_id;
            

            if(record_data.account_id==0){
                selected_account=1;
            }else{
                selected_account=record_data.account_id;
            }
            is_mutable=String(record_data.is_mutable);

        }else{
            selected_category=23;
            selected_account=3;
        }
    });

</script>


<div class="flex-l mx-52 bg-white px-10 pt-5 pb-5 rounded-md">
    <div class="">
        <div class="inline-flex ml-5 mr-5">
            <label for="amount" class="block text-sm font-medium text-gray-700 mr-3">
                Monto
            </label>
            <div class="relative mt-1 rounded-md shadow-sm">
                <div class="absolute inset-y-0 left-0 flex items-center pl-3 pointer-events-none">
                    <span class="text-gray-500 sm:text-sm">
                        $
                    </span>
                </div>
                <input type="text" name="amount" id="amount" class="block w-full px-4 py-2 pr-12 border-t border-b border-l border-gray-300 rounded-md focus:ring-indigo-500 focus:border-indigo-500 pl-7 sm:text-sm" placeholder="0.00" bind:value={amount} autocomplete="off"/>
                <div class="absolute inset-y-0 right-0 flex items-center">
                    <label for="currency" class="sr-only">
                        Currency
                    </label>
                </div>
            </div>
        </div>
        <div class="inline-flex ml-5 mr-5">
            <label for="amount_io" class="block text-sm font-medium text-gray-700 mr-3">
                Tipo
            </label>
            <div class="relative mt-1 rounded-md shadow-sm">
                <select class="block w-full px-4 py-2 pr-12 border-t border-b border-l border-gray-300 rounded-md focus:ring-indigo-500 focus:border-indigo-500 pl-7 sm:text-sm" name="amount_io" bind:value={selected_amount_io}>
                    <option value="out" selected>Egreso</option>
                    <option value="in">Ingreso</option>
                </select>
                <div class="absolute inset-y-0 right-0 flex items-center">
                    <label for="amount_io" class="sr-only">
                        amount_io
                    </label>
                </div>
            </div>
        </div>
        
        <div class="inline-flex ml-5 mr-5">
            <label for="name" class="block text-sm font-medium text-gray-700 mr-3">
                Nombre
            </label>
            <div class="relative mt-1 rounded-md shadow-sm">
                <div class="absolute inset-y-0 left-0 flex items-center pl-3 pointer-events-none">
                    <span class="text-gray-500 sm:text-sm">
                        >
                    </span>
                </div>
                <input type="text" name="name" id="name" class="block w-full px-4 py-2 pr-12 border-t border-b border-l border-gray-300 rounded-md focus:ring-indigo-500 focus:border-indigo-500 pl-7 sm:text-sm" placeholder="Producto / Servicio ..." bind:value={name} autocomplete="off"/>
                <div class="absolute inset-y-0 right-0 flex items-center">
                    <label for="name" class="sr-only">
                        name
                    </label>
                </div>
            </div>
        </div>
    </div>
    
    <div class="mt-8">
        <div class="inline-flex ml-5 mr-5">
            <label for="name" class="block text-sm font-medium text-gray-700 mr-3">
                Commentario
            </label>
            <div class="relative mt-1 rounded-md shadow-sm">
                <div class="absolute inset-y-0 left-0 flex items-center pl-3 pointer-events-none">
                    <span class="text-gray-500 sm:text-sm">
                        >
                    </span>
                </div>
                <input type="text" name="comment" id="comment" class="block w-full px-4 py-2 pr-12 border-t border-b border-l border-gray-300 rounded-md focus:ring-indigo-500 focus:border-indigo-500 pl-7 sm:text-sm" placeholder="Producto / Servicio ..." bind:value={comment} autocomplete="off"/>
                <div class="absolute inset-y-0 right-0 flex items-center">
                    <label for="comment" class="sr-only">
                        comment
                    </label>
                </div>
            </div>
        </div>
    </div>
    <div class="mt-8">
        <div class="inline-flex ml-5 mr-5">
            <label for="record_date" class="block text-sm font-medium text-gray-700 mr-3">
                Fecha
            </label>
            <div class="relative mt-1 rounded-md shadow-sm">
                <input type="text" name="record_date" id="record_date" class="block w-full px-4 py-2 pr-12 border-t border-b border-l border-gray-300 rounded-md focus:ring-indigo-500 focus:border-indigo-500 pl-7 sm:text-sm" placeholder="26/04/2023" bind:value={now_today}/>
                <div class="absolute inset-y-0 right-0 flex items-center">
                    <label for="record_date" class="sr-only">
                        record_date
                    </label>
                </div>
            </div>
        </div>
        <div class="inline-flex ml-5 mr-5">
            <label for="category_id" class="block text-sm font-medium text-gray-700 mr-3">
                Categoria
            </label>
            <div class="relative mt-1 rounded-md shadow-sm">
                <select class="block w-full px-4 py-2 pr-12 border-t border-b border-l border-gray-300 rounded-md focus:ring-indigo-500 focus:border-indigo-500 pl-7 sm:text-sm" name="category_id" bind:value={selected_category}>
                    {#each categories as cat}
                        <option value={cat.ID}>{cat.name}</option>
                    {/each }
                </select>
                <div class="absolute inset-y-0 right-0 flex items-center">
                    <label for="category_id" class="sr-only">
                        category_id
                    </label>
                </div>
            </div>
        </div>
        <div class="inline-flex ml-5 mr-5">
            <label for="account" class="block text-sm font-medium text-gray-700 mr-3">
                Cuenta
            </label>
            <div class="relative mt-1 rounded-md shadow-sm">
                <select class="block w-full px-4 py-2 pr-12 border-t border-b border-l border-gray-300 rounded-md focus:ring-indigo-500 focus:border-indigo-500 pl-7 sm:text-sm" name="account" bind:value={selected_account}>
                    {#each accounts as account}
                        <option value={account.ID}>{account.name}</option>
                    {/each }
                </select>
                <div class="absolute inset-y-0 right-0 flex items-center">
                    <label for="account" class="sr-only">
                        account
                    </label>
                </div>
            </div>
        </div>
        <div class="inline-flex ml-5 mr-5">
            <label for="is_mutable" class="block text-sm font-medium text-gray-700 mr-3">
                Mutable?
            </label>
            <div class="relative mt-1 rounded-md shadow-sm">
                <select class="block w-full px-4 py-2 pr-12 border-t border-b border-l border-gray-300 rounded-md focus:ring-indigo-500 focus:border-indigo-500 pl-7 sm:text-sm" name="is_mutable" bind:value={is_mutable}>
                    <option value="true" selected>Sí</option>
                    <option value="false">No</option>
                </select>
                <div class="absolute inset-y-0 right-0 flex items-center">
                    <label for="is_mutable" class="sr-only">
                        is_mutable
                    </label>
                </div>
            </div>
        </div>
    </div>
    <div class="mt-10">
        <Btnsubmit btnText="Guardar" btnTextOnSubmitting="Guardando..." url={endpoint} method={submitMethod} object={getNewRecord} goback={true}></Btnsubmit>
    </div>
</div>
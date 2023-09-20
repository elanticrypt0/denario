<script lang="ts">

    import { onMount } from 'svelte';
    import apiURL from '../../api/api';
    
    import { ApiCategories } from '../../api/categories';
    import { DateHelper } from '../../helpers/datehelper';
    import { ApiAccounts } from '../../api/accounts';
    import Btnsubmit from '../btnsubmit.svelte';
    


    let now_today=DateHelper.getFullDate();
    
    let selected_account_sender:any=3;
    let selected_account_reciber:any=1;


    let accounts:any[] | any=[];
    let amount:string;
    let submitMethod="POST";
    let endpoint="records/accountsmove/";

    
    async function getAccounts(){
        const response=await fetch(apiURL+"accounts");
        const data=await response.json();
        accounts=data;
    }


    function getAccountsMove():string{
        return JSON.stringify({
            account_sender:Number(selected_account_sender),
            account_reciber:Number(selected_account_reciber),
            amount:Number(amount),
            record_date:now_today,
        });
    }
    

    onMount(async () =>{
        const[year,month]=DateHelper.getCurrentYearAndMonth();
        endpoint+=year+"/"+month;
        
        await getAccounts();
            
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
            <label for="account" class="block text-sm font-medium text-gray-700 mr-3">
                Cuenta de origen
            </label>
            <div class="relative mt-1 rounded-md shadow-sm">
                <select class="block w-full px-4 py-2 pr-12 border-t border-b border-l border-gray-300 rounded-md focus:ring-indigo-500 focus:border-indigo-500 pl-7 sm:text-sm" name="account" bind:value={selected_account_sender}>
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
            <label for="account" class="block text-sm font-medium text-gray-700 mr-3">
                Cuenta de destino
            </label>
            <div class="relative mt-1 rounded-md shadow-sm">
                <select class="block w-full px-4 py-2 pr-12 border-t border-b border-l border-gray-300 rounded-md focus:ring-indigo-500 focus:border-indigo-500 pl-7 sm:text-sm" name="account" bind:value={selected_account_reciber}>
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

    </div>
    <div class="mt-10">
        <Btnsubmit btnText="Transferir" btnTextOnSubmitting="Transfiriendo..." url={endpoint} method={submitMethod} object={getAccountsMove}></Btnsubmit>
    </div>
</div>
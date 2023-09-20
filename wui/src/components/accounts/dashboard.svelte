<script lang="ts">

    import { onMount } from 'svelte';
    
    import apiURL from '../../api/api';
    import type Record from '../../interfaces/record';
    import {ApiRecords} from '../../api/records';
    import { DateHelper } from '../../helpers/datehelper';
    
    
    // dates
    let current_year=2022;
    let current_month="01";
    
    // records
    let accounts_list:any[] | any=[];    
      
    async function getAccounts(year:number,month:string){
        const response=await fetch(apiURL+"accounts");
        const data=await response.json();
        accounts_list=data;
    }
     
    onMount(async () =>{
        [current_year,current_month]=DateHelper.getCurrentYearAndMonth();
        getAccounts(current_year,current_month);
    });
    
    
    </script>
    
   
    <div class="flex justify-end pt-5">
        <div class="pr-2">
        <a href="/accounts/create" class="py-2 px-4 flex justify-center items-center  bg-green-500 hover:bg-green-700 focus:ring-green-500 focus:ring-offset-green-200 text-white w-full transition ease-in duration-200 text-center text-base font-semibold shadow-md focus:outline-none focus:ring-2 focus:ring-offset-2  rounded-full">
            Nueva cuenta
        </a>
        </div>
        <div class="pl-5">
            <input type="text" class="rounded-lg border-transparent flex-1 appearance-none border border-gray-300 w-full py-2 px-4 bg-white text-gray-700 placeholder-gray-400 shadow-sm text-base focus:outline-none focus:ring-2 focus:ring-purple-600 focus:border-transparent" placeholder="Buscar" />
        </div>
    </div>
    
    <div class="relative overflow-x-auto shadow-md sm:rounded-lg mt-5">
        <table class="w-full text-sm text-left text-gray-500 dark:text-gray-400">
            <thead class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
                <tr>
                    <th scope="col" class="p-4">
                        <div class="flex items-center">
                            <input id="checkbox-all-search" type="checkbox" class="w-4 h-4 text-blue-600 bg-gray-100 border-gray-300 rounded focus:ring-blue-500 dark:focus:ring-blue-600 dark:ring-offset-gray-800 dark:focus:ring-offset-gray-800 focus:ring-2 dark:bg-gray-700 dark:border-gray-600">
                            <label for="checkbox-all-search" class="sr-only">checkbox</label>
                        </div>
                    </th>
                    <th scope="col" class="px-6 py-3">
                        Nombre
                    </th>
                    <th scope="col" class="px-6 py-3">
                        Short
                    </th>
                    <th scope="col" class="px-6 py-3">
                        Detalle
                    </th>
                    <th scope="col" class="px-6 py-3">
                        Ax
                    </th>
                </tr>
            <tbody>
                {#each accounts_list as account }
                <tr class="bg-white border-b dark:bg-gray-800 dark:border-gray-700 hover:bg-gray-50 dark:hover:bg-gray-600">
                    
                    <td class="w-4 p-4">
                        <div class="flex items-center">
                            <input id="checkbox-table-search-1" type="checkbox" class="w-4 h-4 text-blue-600 bg-gray-100 border-gray-300 rounded focus:ring-blue-500 dark:focus:ring-blue-600 dark:ring-offset-gray-800 dark:focus:ring-offset-gray-800 focus:ring-2 dark:bg-gray-700 dark:border-gray-600">
                            <label for="checkbox-table-search-1" class="sr-only">checkbox</label>
                        </div>
                    </td>
                    <th scope="row" class="px-6 py-4 font-medium text-gray-900 whitespace-nowrap dark:text-white">
                        {account.name}
                    </th>
                    <td class="px-6 py-4">
                        {account.short}
                    </td>
                    <td class="px-6 py-4">
                        {account.detail}
                    </td>
                    <td class="flex items-center px-6 py-4 space-x-3">
                        <a href="/accounts/edit/#{account.ID}" class="font-medium text-blue-600 dark:text-blue-500 hover:underline">Edit</a>                    
                        <a href="/accounts/delete/#{account.ID}" class="font-medium text-red-600 dark:text-red-500 hover:underline">Remove</a>
                    </td>
                </tr>
                {/each}
            </tbody>
        </table>
    </div>
    
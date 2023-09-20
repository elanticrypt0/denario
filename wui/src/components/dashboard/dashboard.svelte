<script lang="ts">

import { onMount } from 'svelte';
import Balances from './balances.svelte';

import apiURL from '../../api/api';
import type Record from '../../interfaces/record';
import {ApiRecords} from '../../api/records';
import { DateHelper } from '../../helpers/datehelper';


// dates
let current_year=2022;
let current_month="01";

// records
let records_list:any[] | any=[];

// para buscar
let records_list_original:any[]=[];
let search_text="";

// balances
let ingresos=1000;
let egresos=500;
let fijos=10;
let creditos=12.50;
let ahorros=1.00;
// accounts
let account_bp=1.00;
let account_efe=1.00;
let account_mp=1.00;


// total amount
let records_selected:any[]=[];
let total_amount=0;

// $: sumAmount=()=>{
//     const initialValue = 0;
//     total_amount = records_selected.reduce((accumulator, currentValue) => accumulator + currentValue, initialValue);
// }

$:{
    sumAmounts(records_selected);
}

function sumAmounts(records_selected: any[]){
    const initialValue = 0;
    total_amount = records_selected.reduce((accumulator, currentValue) => {
        currentValue=Number(currentValue.substring(currentValue.indexOf("-")+1));
        return accumulator + currentValue
    }, initialValue);
}

// deschequear los checks
function uncheckSumAmounts(){
    total_amount=0;
    records_selected=[];
}

function goToToday(){
    [current_year,current_month]=DateHelper.getCurrentYearAndMonth();
    getRecords(current_year,current_month);
    getBalances(current_year,current_month);
    uncheckSumAmounts();
}

function goNextMonth(){
    [current_month,current_year]=DateHelper.getNextMonth(current_month,current_year);
    getRecords(current_year,current_month);
    uncheckSumAmounts();
    getBalances(current_year,current_month);
}

function goPrevMonth(){
    [current_month,current_year]=DateHelper.getPrevMonth(current_month,current_year);
    getRecords(current_year,current_month);
    getBalances(current_year,current_month);
    uncheckSumAmounts();
}

async function getRecords(year:number,month:string){
    records_list=[];
    const response=await fetch(apiURL+"records/"+year+"/"+month);
    const data=await response.json();
    records_list_original=data;
    records_list=records_list_original;

}

function deleteRecord(id:Number):null{
    // console.log(id);
    const url="records/"+id;
    // console.log(apiURL+url);
    const response= fetch(apiURL+url,{
        method:'DELETE',
        headers:{
            'Content-Type':'application/json'
        }
    });
    return null;
}

async function getBalancesInOut(year:number,month:string){

    const response=await fetch(apiURL+"balances/out/"+year+"/"+month+"/fixed");
    const data=await response.json();
    fijos=data.total;

    const response2=await fetch(apiURL+"balances/out/"+year+"/"+month+"/credits");
    const data2=await response2.json();
    creditos=data2.total;
}

async function getBalancesCreditsFixed(year:number,month:string){
    const response=await fetch(apiURL+"balances/in/"+year+"/"+month);
    const data=await response.json();
    ingresos=data.total;

    const response2=await fetch(apiURL+"balances/out/"+year+"/"+month);
    const data2=await response2.json();
    egresos=data2.total;
}

async function getBalancesAccounts(year:number,month:string){
    const response=await fetch(apiURL+"balances/account/3/"+year+"/"+month);
    const data=await response.json();
    account_bp=data.total;

    console.log(apiURL+"balances/account/1/"+year+"/"+month);
    const response2=await fetch(apiURL+"balances/account/1/"+year+"/"+month);
    const data2=await response2.json();
    account_efe=data2.total;

    const response3=await fetch(apiURL+"balances/account/2/"+year+"/"+month);
    const data3=await response3.json();account_mp=data3.total;
}

async function getBalances(year:number,month:string){
    getBalancesInOut(year,month);
    getBalancesCreditsFixed(year,month);
    getBalancesAccounts(year,month);
}

function search(){
    // if(search_text!=""){
    //     console.log(search_text);
    //     const search_regex=new RegExp(`${search_text}[a-z]*`);
    //     console.log(search_regex);
    //     // busca dentro de records_list_original y elimina todo lo que no coincide con la búsqueda
    //     records_list=[];
    //     records_list=records_list_original.map((el)=>{
    //         // console.log(el);
    //         console.log(el.name.match(search_regex));
    //         // return el.name.match(search_regex);
    //         if()
    //     });      
    // }else{
    //     records_list=records_list_original;
    // }
}

onMount(async () =>{
    [current_year,current_month]=DateHelper.getCurrentYearAndMonth();
    getRecords(current_year,current_month);
    getBalances(current_year,current_month);
});


</script>

<Balances ingresos={ingresos} egresos={egresos} fijos={fijos} creditos={creditos} ahorros={ahorros} account_bp={account_bp} account_efe={account_efe} account_mp={account_mp}></Balances>

<div class="flex justify-end pt-5">
    <div class="pr-2">
        <button class="py-2 px-4  bg-indigo-600 hover:bg-indigo-700 focus:ring-indigo-500 focus:ring-offset-indigo-200 text-white w-full transition ease-in duration-200 text-center text-base font-semibold shadow-md focus:outline-none focus:ring-2 focus:ring-offset-2  rounded-lg" on:click={goPrevMonth}>Prev.</button>
    </div>
    <div class="pr-2">
        <button>{current_year} / {current_month}</button>
    </div>
    <div class="pr-2">
        <button class="py-2 px-4  bg-indigo-600 hover:bg-indigo-700 focus:ring-indigo-500 focus:ring-offset-indigo-200 text-white w-full transition ease-in duration-200 text-center text-base font-semibold shadow-md focus:outline-none focus:ring-2 focus:ring-offset-2  rounded-lg" on:click={goNextMonth}>Prox.</button>
    </div>
    <div class="pr-20">
        <button class="py-2 px-4  bg-indigo-600 hover:bg-indigo-700 focus:ring-indigo-500 focus:ring-offset-indigo-200 text-white w-full transition ease-in duration-200 text-center text-base font-semibold shadow-md focus:outline-none focus:ring-2 focus:ring-offset-2  rounded-lg" on:click={goToToday}>Hoy</button>
    </div>
    <div class="pr-2">

        <div class="inline-flex ml-5 mr-5">
            <label for="amount" class="block text-sm font-medium text-gray-700 mr-3">
                Suma total
            </label>
            <div class="relative mt-1 rounded-md shadow-sm">
                <div class="absolute inset-y-0 left-0 flex items-center pl-3 pointer-events-none">
                    <span class="text-gray-500 sm:text-sm">
                        $
                    </span>
                </div>
                <input type="text" name="amount" id="amount" class="block w-full px-4 py-2 pr-12 border-t border-b border-l border-gray-300 rounded-md focus:ring-indigo-500 focus:border-indigo-500 pl-7 sm:text-sm" placeholder="0.00" bind:value={total_amount} autocomplete="off"/>
                <div class="absolute inset-y-0 right-0 flex items-center">
                    <label for="currency" class="sr-only">
                        Currency
                    </label>
                </div>
            </div>
        </div>
    </div>
    <div class="pr-2">
    <a href="/records/create" class="py-2 px-4 flex justify-center items-center  bg-green-500 hover:bg-green-700 focus:ring-green-500 focus:ring-offset-green-200 text-white w-full transition ease-in duration-200 text-center text-base font-semibold shadow-md focus:outline-none focus:ring-2 focus:ring-offset-2  rounded-full">
        Registrar
    </a>
    </div>
    <div class="pr-2">
    <a href="/records/accountsmove" class="py-2 px-4 flex justify-center items-center  bg-violet-400 hover:bg-violet-700 focus:ring-green-500 focus:ring-offset-green-200 text-white w-full transition ease-in duration-200 text-center text-base font-semibold shadow-md focus:outline-none focus:ring-2 focus:ring-offset-2  rounded-full">
        Transferir
    </a>

    </div>
    <div class="pl-5">
        <input type="text" class="rounded-lg border-transparent flex-1 appearance-none border border-gray-300 w-full py-2 px-4 bg-white text-gray-700 placeholder-gray-400 shadow-sm text-base focus:outline-none focus:ring-2 focus:ring-purple-600 focus:border-transparent" bind:value={search_text} on:change={search} placeholder="Buscar" />
    </div>
</div>

<div class="relative overflow-x-auto shadow-md sm:rounded-lg mt-5">
    <table class="w-full text-sm text-left text-gray-500 dark:text-gray-400">
        <thead class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
            <tr>
                <th scope="col" class="p-4">
                    <div class="flex items-center">
                        <!-- <input id="checkbox-all-search" type="checkbox" class="w-4 h-4 text-blue-600 bg-gray-100 border-gray-300 rounded focus:ring-blue-500 dark:focus:ring-blue-600 dark:ring-offset-gray-800 dark:focus:ring-offset-gray-800 focus:ring-2 dark:bg-gray-700 dark:border-gray-600">
                        <label for="checkbox-all-search" class="sr-only">checkbox</label> -->
                        <button type="button" class="w-full text-base px-1 rounded-md text-gray-600 bg-white border hover:bg-gray-100" on:click={uncheckSumAmounts}>
                            U
                        </button>
                    </div>
                </th>
                <th scope="col" class="">
                    I/O
                </th>
                <th scope="col" class="px-6 py-3">
                    Nombre
                </th>
                <th scope="col" class="px-6 py-3">
                    Monto
                </th>
                <th scope="col" class="px-6 py-3">
                    Cat.
                </th>
                <th scope="col" class="px-6 py-3">
                    Cuenta
                </th>
                <th scope="col" class="px-6 py-3">
                    Detalle
                </th>
                <th scope="col" class="px-6 py-3">
                    Fecha
                </th>
                <th scope="col" class="px-6 py-3">
                    Ax
                </th>
            </tr>
        <tbody>
            {#each records_list as record }
            <tr class="bg-white border-b dark:bg-gray-800 dark:border-gray-700 hover:bg-gray-50 dark:hover:bg-gray-600">
                
                <td class="w-4 p-4">
                    <div class="flex items-center">
                        <input id="checkbox-table-search-1" type="checkbox" value={record.ID+"-"+record.amount} bind:group={records_selected} class="w-4 h-4 text-blue-600 bg-gray-100 border-gray-300 rounded focus:ring-blue-500 dark:focus:ring-blue-600 dark:ring-offset-gray-800 dark:focus:ring-offset-gray-800 focus:ring-2 dark:bg-gray-700 dark:border-gray-600">
                        <label for="checkbox-table-search-1" class="sr-only">checkbox</label>
                    </div>
                </td>
                {#if record.amount_io=="out"}
                    <td class="px-1 py-1 bg-red-300">&nbsp;</td>
                {:else}
                    <td class="px-1 py-1 bg-green-300">&nbsp;</td>
                {/if}
                <th scope="row" class="px-6 py-4 font-medium text-gray-900 whitespace-nowrap dark:text-white">
                    {record.name}
                </th>
               
                <td class="px-6 py-4">
                    ${record.amount}
                </td>
                <td class="px-6 py-4">
                    {#if record.Category.ID==28}
                    <a class="px-2 py-1  text-xs rounded text-white  bg-yellow-400 font-medium" href="#{record.Category.ID}">{record.Category.name}</a>
                    {:else if record.Category.ID==15}
                    <a class="px-2 py-1  text-xs rounded text-white  bg-teal-400 font-medium" href="#{record.Category.ID}">{record.Category.name}</a>
                    {:else if record.Category.ID > 8 && record.Category.ID <14}
                    <a class="px-2 py-1  text-xs rounded text-white  bg-pink-400 font-medium" href="#{record.Category.ID}">{record.Category.name}</a>
                    {:else if record.Category.ID == 14}
                    <a class="px-2 py-1  text-xs rounded text-white  bg-green-500 font-medium" href="#{record.Category.ID}">{record.Category.name}</a>
                    {:else}
                    <a class="px-2 py-1  text-xs rounded text-white  bg-blue-400 font-medium" href="#{record.Category.ID}">{record.Category.name}</a>
                    {/if}
                </td>
                <td class="px-6 py-4">
                    {#if record.Category.ID!=28}
                        {#if record.Account.ID==3}
                        <a class="px-2 py-1  text-xs rounded text-white  bg-lime-600 font-medium" href="#{record.Account.ID}">{record.Account.short}</a>
                        {:else if record.Account.ID==2}
                        <a class="px-2 py-1  text-xs rounded text-white  bg-blue-700 font-medium" href="#{record.Account.ID}">{record.Account.short}</a>
                        {:else if record.Account.ID==1}
                        <a class="px-2 py-1  text-xs rounded text-white  bg-green-600 font-medium" href="#{record.Account.ID}">{record.Account.short}</a>
                        {:else}
                        <a class="px-2 py-1  text-xs rounded text-white  bg-slate-400 font-medium" href="#{record.Account.ID}">{record.Account.short}</a>
                        {/if}
                    {/if}
                </td>
                <td class="px-6 py-4">
                    {record.comment}
                </td>
                <td class="px-6 py-4">
                    {record.record_date}
                </td>
                <td class="flex items-center px-6 py-4 space-x-3">
                    {#if record.is_mutable==true}
                    <a href="/records/edit/#{record.ID}" class="font-medium text-blue-600 dark:text-blue-500 hover:underline">Edit</a>                    
                    <a href="#delete/{record.ID}" class="font-medium text-red-600 dark:text-red-500 hover:underline" on:click={deleteRecord(record.ID)}>Remove</a>
                    {:else}
                    -
                    {/if}
                </td>
            </tr>
            {/each}
        </tbody>
    </table>
</div>

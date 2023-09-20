<script>
    export let ingresos=0.00;
    export let egresos=0.00;
    export let fijos=0.00;
    export let creditos=0.00;
    export let ahorros=0.00;

    export let account_bp=0.00;
    export let account_efe=0.00;
    export let account_mp=0.00;

</script>
<div class="gap-4 flex justify-center">
    <span class="px-4 py-2 text-sm rounded-full text-slate-950 border border-green-400">Ingresos ${ingresos}</span>
    <span class="px-4 py-2 text-sm rounded-full text-slate-950 border border-red-400">Egresos ${egresos}</span>
    <span class="px-4 py-2 text-sm rounded-full text-slate-950 border border-yellow-400">Diferencias ${ingresos - egresos}</span>
    
    <span class="px-4 py-2  text-sm rounded-full text-black-500 border border-black-400">|</span>

    <span class="px-4 py-2  text-sm rounded-full text-slate-950 border border-indigo-400">Fijos ${fijos}</span>
    <span class="px-4 py-2  text-sm rounded-full text-slate-950 border border-purple-400">Créditos ${creditos}</span>
    <span class="px-4 py-2  text-sm rounded-full text-slate-950 border border-blue-400">Fijos + créditos ${fijos + creditos}</span>
    <span class="px-4 py-2  text-sm rounded-full text-slate-950 border border-pink-400">Ahorros ${ahorros}</span>
    
    <span class="px-4 py-2  text-sm rounded-full text-black-500 border border-black-400">|</span>

    <span class="px-4 py-2  text-sm rounded-full text-slate-950 border border-green-400">Cuenta BP ${account_bp}</span>
    <span class="px-4 py-2  text-sm rounded-full text-slate-950 border border-blue-400">Cuenta EFE ${account_efe}</span>
    <span class="px-4 py-2  text-sm rounded-full text-slate-950 border border-yellow-400">Cuenta MP ${account_mp}</span>
    <span class="px-4 py-2  text-sm rounded-full text-slate-950 border border-pink-400">Total ${account_bp+account_efe+account_mp}</span>
</div>
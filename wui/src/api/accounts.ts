import {apiURL} from './api';

export class ApiAccounts{

    static async get4Select():any{
        const response=await fetch(apiURL+"accounts");
        const data=await response.json();
        let accounts=[];
        data.forEach(elem => {
            accounts.push({'value':elem.ID,'name':elem.name});
        });
        return accounts;
    }

    static async getAll():Promise<Response>{

        const result=await fetch(apiURL+"accounts/",{
            method:'GET',
            mode:'cors',
            headers:{},
        })
        .then((result) => {
        return result.json();
        });

        return result;
    }

    static async create(name:string,amount:number,amount_io:string,comment:string,record_date:string,category_id:number,is_mutable:boolean):Promise<Response>{

        const data={
            name,        
            amount,
            amount_io,
            comment,
            record_date,
            category_id,
            is_mutable
        }

        const result=await fetch(apiURL+"accounts/",{
            method:'POST',
            mode:'cors',
            headers:{},
            body:JSON.stringify(data),
        })
        .then((result) => {
        return result.json();
        });

        return result;
    }

    static async update(id:number,name:string,amount:number,amount_io:string,comment:string,record_date:string,category_id:number,is_mutable:boolean):Promise<Response>{

        const data={
            name,        
            amount,
            amount_io,
            comment,
            record_date,
            category_id,
            is_mutable
        }

        const result=await fetch(apiURL+"accounts/"+id,{
            method:'PUT',
            mode:'cors',
            headers:{},
            body:JSON.stringify(data),
        })
        .then((result) => {
        return result.json();
        });

        return result;
    }

    static async delete(id:number):Promise<Response>{

        const result=await fetch(apiURL+"accounts/"+id,{
            method:'DELETE',
            mode:'cors',
            headers:{},
        })
        .then((result) => {
        return result.json();
        });

        return result;
    }
}
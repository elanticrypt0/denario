import {apiURL} from './api';

export class ApiFixed{
    static async getAll(year:string,month:string):Promise<Response>{

        const data={
            year,
            month
        }
    
        const result=await fetch(apiURL+"credits/"+year+"/"+month,{
            method:'GET',
            mode:'cors',
            headers:{},
            body:JSON.stringify(data),
        })
        .then((result) => {
           return result.json();
        });
    
        return result;
    }
    
    static async getOne(id:number):Promise<Response>{
    
        const result=await fetch(apiURL+"credits/"+id,{
            method:'GET',
            mode:'cors',
            headers:{},
        })
        .then((result) => {
           return result.json();
        });
    
        return result;
    }
    
    static async create(name:string,comment:string,amount:number,payments:number,started_at:string,category_id:number,is:boolean):Promise<Response>{
    
        const data={
            name,        
            comment,
            amount,
            payments,
            started_at,
            category_id,
            is
        }
    
        const result=await fetch(apiURL+"credits/",{
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
    
    static async update(id:number,name:string,comment:string,amount:number,payments:number,started_at:string,category_id:number,is:boolean):Promise<Response>{
    
        const data={
            name,        
            comment,
            amount,
            payments,
            started_at,
            category_id,
            is
        }
    
        const result=await fetch(apiURL+"credits/"+id,{
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
    
        const result=await fetch(apiURL+"credits/"+id,{
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
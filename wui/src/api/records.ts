import apiURL from "./api";


export class ApiRecords{

    static async getRecordsOfYearMonth(year:string,month:string):Promise<Response>{

        console.log(apiURL+"records/"+year+"/"+month);
        const result=await fetch(apiURL+"records/"+year+"/"+month,{
            method:'GET',
            mode:'cors',
            headers:{},
        })
        .then((result) => {
            return result;
        });
        const data=result.json()
        return data;
    }

    static async create(name:string,amount:number,amount_io:string,comment:string,record_date:string,category_id:number,is_mutable:boolean,account_id:number):Promise<Response>{

        const data={
            name,        
            amount,
            amount_io,
            comment,
            record_date,
            category_id,
            is_mutable,
            account_id
        }

        const result=await fetch(apiURL+"records/",{
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

        const result=await fetch(apiURL+"records/"+id,{
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

        const result=await fetch(apiURL+"records/"+id,{
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
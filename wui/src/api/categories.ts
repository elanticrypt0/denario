import {apiURL} from './api';

export class ApiCategories{

    static async get4Select():Promise<any>{
        const response=await fetch(apiURL+"categories");
        const data=await response.json();
        let cats=[];
        data.forEach(elem => {
            cats.push({'value':elem.ID,'name':elem.name});
        });
        return cats;
    }

    static async getAll(year:string,month:string):Promise<Response>{

        const data={
            year,
            month
        }
    
        const result=await fetch(apiURL+"categories/"+year+"/"+month,{
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
    
        const result=await fetch(apiURL+"categories/"+id,{
            method:'GET',
            mode:'cors',
            headers:{},
        })
        .then((result) => {
           return result.json();
        });
    
        return result;
    }
    
    static async create(name:string):Promise<Response>{
    
        const data={
            name,
        }
    
        const result=await fetch(apiURL+"categories/",{
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
    
    static async update(id:number,name:string):Promise<Response>{
    
        const data={
            name,        
        }
    
        const result=await fetch(apiURL+"categories/"+id,{
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
    
        const result=await fetch(apiURL+"categories/"+id,{
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
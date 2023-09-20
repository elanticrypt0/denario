export class DateHelper{


    static getCurrentYearAndMonth():[number,string]{
        const year=Number(new Date().getFullYear());
        let month=String(new Date().getMonth()+1);
        month=this.addZero2String(month);
        return [year,month];
    }

    static getCurrentDay():string{
        const day=Number(new Date().getDate());
        const current_day=this.addZeroToNumber(day);
        return current_day;
    }
    
    static getCurrentDate():string{
        const [year,month]=this.getCurrentYearAndMonth();
        const day=this.getCurrentDay();
        return `${year}-${month}-${day}`;

    }

    static getFullDate():string{
        const date=this.getCurrentDate()
        const hours=this.addZeroToNumber(new Date().getHours());
        const min=this.addZeroToNumber(new Date().getMinutes());
        const sec=this.addZeroToNumber(new Date().getSeconds());
        return `${date} ${hours}:${min}:${sec}`;
    }


    static getPrevMonth(month:string,year:number):[string,number]{
        let month_aux=Number(month);
    
        if(month_aux-1 == 0){
            month=String(12);
            year=year-1;
            return [month,year];
        }
    
        month=String(month_aux-1);
        month=this.addZero2String(month);
        return [month,year];
    }
    
    static getNextMonth(month:string,year:number):[string,number]{
        let month_aux=Number(month);
        
        if(month_aux+1 == 13){
            month=String(1);
            year=year+1;
            month=this.addZero2String(month);
            return [month,year];
        }
    
        month=String(month_aux+1);
        month=this.addZero2String(month);
        return [month,year];
    }

    static addZero2String(month:string):string{
        if(Number(month) < 10){
            month="0"+month;
        }
        return month
    }

    static addZeroToNumber(num:number):string{
        let num_string="0";
        if(num<10){
            num_string+=num;
        }else{
            num_string=String(num);
        }
        return num_string;
    }


}
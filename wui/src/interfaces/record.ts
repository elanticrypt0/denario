import type { Category } from "./category";

interface Record {
  ID: number;
  CreatedAt: string;
  UpdatedAt: string;
  DeletedAt?: any;
  name: string;
  amount: number;
  amount_io: string;
  comment: string;
  record_date: string;
  category_id: number;
  Category: Category;
  is_mutable: boolean;
  account_id:number,
  Account:any
}

export default Record;
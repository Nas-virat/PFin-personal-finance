import { fetchdata } from "@/utils/Fetch";



const createTransaction = async (
    transaction_type : string,
    category: string,
    description: string,
    amount: number,
) => {
    try{
        const res = await fetchdata.post('/transaction/create', {
            transaction_type: transaction_type,
            category: category,
            description: description,
            amount: amount,
        });
        return res.data;
    }
    catch(error){
        console.error('Error creating transaction:', error);
        throw error;
    }
}

const getTransactionsByMonthYear = async (month: number, year: number) => {
    try{
        const res = await fetchdata.get(`/transaction/month/${month}/year/${year}`);
        return res.data;
    }
    catch(error){
        console.error('Error getting transactions:', error);
        throw error;
    }
}

const getTransactionsByDayMonthYear = async (day:number, month: number, year: number) => {
    try{
        const res = await fetchdata.get(`/transaction/day/${day}/month/${month}/year/${year}`);
        return res.data;
    } catch(error){
        console.error('Error getting transactions:', error);
        throw error;
    }
}
    
export {    
        createTransaction,
        getTransactionsByMonthYear ,
        getTransactionsByDayMonthYear,
    };
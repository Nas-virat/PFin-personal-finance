import { fetchdata } from "@/utils/Fetch";

const getSummaryBalance = async () =>{
    try{
        const res = await fetchdata.get('/balance/summary');
        return res.data;
    }
    catch(error){
        console.error('Error getting transactions:', error);
        throw error;
    }
}


export {getSummaryBalance};
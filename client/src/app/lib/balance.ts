import { fetchdata } from "@/utils/Fetch";

const getSummaryBalance = async () =>{
    try{
        const res = await fetch(`http://localhost:8000/api/balance/summary`, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
            },
        });
        if (!res.ok) {
            throw new Error(`Failed to get transactions. Status: ${res.status}`);
        }

        const data = await res.json();
        return data;
    }
    catch(error){
        console.error('Error getting transactions:', error);
        throw error;
    }
}


export {getSummaryBalance};
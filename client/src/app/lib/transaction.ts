


const createTransaction = async (
    transaction_type : string,
    category: string,
    description: string,
    amount: number,
) => {
    try{
        const res = await fetch('http://localhost:8000/api/transaction/create', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                "transaction_type": transaction_type,
                "category":category,
                "amount":amount,
                "description":description,
            }),
        });
        if (!res.ok) {
            throw new Error(`Failed to create account. Status: ${res.status}`);
        }

        const data = await res.json();
        return data;
    }
    catch(error){
        console.error('Error creating transaction:', error);
        throw error;
    }
}

const getTransactionsByMonthYear = async (month: number, year: number) => {
    try{
        const res = await fetch(`http://localhost:8000/api/transaction/month/${month}/year/${year}`, {
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

const getTransactionsByDayMonthYear = async (day:number, month: number, year: number) => {
    try{
        const res = await fetch(`http://localhost:8000/api/transaction/day/${day}/month/${month}/year/${year}`, {
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
    
export {    
        createTransaction,
        getTransactionsByMonthYear ,
        getTransactionsByDayMonthYear,
    };
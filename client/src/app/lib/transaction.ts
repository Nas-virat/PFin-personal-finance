

const createTransaction = async (
    transaction_type : string,
    catergory: string,
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
                "catergory":catergory,
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
    
export { createTransaction };
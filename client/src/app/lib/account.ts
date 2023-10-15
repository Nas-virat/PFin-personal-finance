

export async function postCreateAccount(
    account_name: string,
    description: string,
    account_type: string,
    amount: number,
) {
    try {
        const res = await fetch('http://localhost:8000/api/account/create', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                "account_name": account_name,
                "account_type":account_type,
                "amount":amount,
                "description":description,
                "currency":"THB"
            }),
        });

        if (!res.ok) {
            throw new Error(`Failed to create account. Status: ${res.status}`);
        }

        const data = await res.json();
        return data;
    } catch (error) {
        console.error('Error creating account:', error);
        throw error;
    }
}
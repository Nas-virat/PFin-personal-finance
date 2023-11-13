import { fetchdata } from "@/utils/Fetch";


const postCreateAccount = async (
    account_name: string,
    description: string,
    account_type: string,
    amount: number,
) => {
    try {
        const res = await fetchdata.post('/account/create', {
            account_name: account_name,
            description: description,
            account_type: account_type,
            amount: amount,
            currency: "THB"
        });
        return res.data;
    } catch (error) {
        console.error('Error creating account:', error);
        throw error;
    }
}

const EditAccountInfo = async (
    account_id: string,
    account_name: string,
    description: string,
    account_type: string,
    amount: number,
) => {
    try {
        const res = await fetchdata.put('/account/edit/' + account_id, {
            account_name: account_name,
            description: description,
            account_type: account_type,
            amount: amount,
            currency: "THB"
        });
        return res.data;
    } catch (error) {
        console.error('Error editing account:', error);
        throw error;
    }
}

export { postCreateAccount, EditAccountInfo};
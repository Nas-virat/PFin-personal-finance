"use client";
import React, { useState } from "react";
import TextField from "@mui/material/TextField";
import Box from "@mui/material/Box";
import MenuItem from "@mui/material/MenuItem";
import Select, { SelectChangeEvent } from "@mui/material/Select";
import { FormControl, InputLabel } from "@mui/material";
import { createTransaction } from "../../lib/transaction";
import Toast from "@/components/Alert";
import Swal from "sweetalert2";
import { useRouter, useSearchParams } from "next/navigation";
import dayjs from "dayjs";

export default function Page() {
  const [transactionType, setType] = useState("income");
  const [category, setCategory] = useState("");
  const [description, setDescription] = useState("");
  const [balance, setBalance] = useState(0);

  const router = useRouter();
  const searchParams = useSearchParams();

  const handleChange = (event: SelectChangeEvent) => {
    setType(event.target.value as string);
    setCategory(""); // Reset category when changing the transaction type.
  };

  const handleCategory = (event: SelectChangeEvent) => {
    setCategory(event.target.value as string);
  };

  const handleSubmit = async (e: React.MouseEvent<HTMLButtonElement>) => {
    e.preventDefault();

    const res = await createTransaction(
      transactionType,
      category,
      description,
      balance,
      parseInt(searchParams.get("date") ?? dayjs().date().toString()),
      parseInt(searchParams.get("month") ?? (dayjs().month() + 1).toString()),
      parseInt(searchParams.get("year") ?? dayjs().year().toString()),
    );
    if (res.success === true) {
      Toast.fire({
        icon: "success",
        title: "Create Transaction Successfully",
      });
      router.push("/home");
    } else if (res.success === false) {
      Swal.fire({
        icon: "error",
        title: "Oops...",
        text: res.data.message,
      });
    }
  };

  const categories = [
    { type: "income", name: "Salary" },
    { type: "income", name: "Bonus" },
    { type: "income", name: "Rent" },
    { type: "expense", name: "Food" },
    { type: "expense", name: "Travel" },
    { type: "expense", name: "Internet" },
    { type: "expense", name: "Electricity" },
    { type: "expense", name: "Water" },
    { type: "credit", name: "Credit Card" },
    { type: "other", name: "Other" },
  ];

  const filteredCategories = categories.filter(
    (item) => item.type === transactionType,
  );

  return (
    <div>
      <div className="flex justify-between mx-32 text-pf-gray-900 font-bold text-2xl">
        <h1 className="text-[48px]">Add Transaction</h1>
      </div>
      <div className="w-full flex justify-center">
        <Box
          sx={{
            width: 700,
            maxWidth: "100%",
          }}
        >
          <FormControl
            fullWidth
            variant="outlined"
            sx={{ marginBottom: "10px", marginTop: "20px" }}
          >
            <InputLabel id="transactionTypeLabel">Transaction</InputLabel>
            <Select
              labelId="transactionTypeLabel"
              id="transactiontype"
              label="transactiontype"
              value={transactionType}
              onChange={handleChange}
            >
              <MenuItem value={"income"}>Income</MenuItem>
              <MenuItem value={"expense"}>Expense</MenuItem>
              <MenuItem value={"other"}>Other</MenuItem>
            </Select>
          </FormControl>
          <FormControl
            fullWidth
            variant="outlined"
            sx={{ marginBottom: "30px", marginTop: "20px" }}
          >
            <InputLabel id="categoryTypeLabel">Category</InputLabel>
            <Select
              labelId="categoryTypeLabel"
              id="categorytype"
              label="categorytype"
              value={category}
              onChange={handleCategory}
            >
              {filteredCategories.map((item, index) => (
                <MenuItem key={index} value={item.name}>
                  {item.name}
                </MenuItem>
              ))}
            </Select>
          </FormControl>
          <TextField
            fullWidth
            label="Description"
            id="Description"
            variant="outlined"
            sx={{ marginBottom: "30px" }}
            value={description}
            onChange={(e) => setDescription(e.target.value)}
          />
          <FormControl
            fullWidth
            variant="outlined"
            sx={{ marginBottom: "30px" }}
          >
            <TextField
              fullWidth
              id="AmountTransaction"
              label="Amount"
              type="number"
              variant="outlined"
              value={balance}
              onChange={(e) => setBalance(Number(e.target.value))}
              InputLabelProps={{
                shrink: true,
              }}
            />
          </FormControl>
          <button
            className="w-full bg-pf-secondary-2 h-16 font-bold text-white text-xl rounded-lg hover:bg-lime-600"
            onClick={(e) => handleSubmit(e)}
          >
            Confirm
          </button>
        </Box>
      </div>
    </div>
  );
}

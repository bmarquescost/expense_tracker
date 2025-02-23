// frontend/src/pages/Home.js
import React, { useState } from "react";
import TicketForm from "../components/TicketForm";

const Home = () => {
  const [tickets, setTickets] = useState([]);
  const [showForm, setShowForm] = useState(false);

  const addTicket = (ticket) => {
    setTickets([...tickets, ticket]);
    setShowForm(false);
  };

  return (
    <div className="flex flex-col items-center justify-center min-h-screen bg-gray-100 p-4">
      {showForm ? (
        <ExpenseForm addExpense={addExpense} />
      ) : (
        <Button onClick={() => setShowForm(true)} className="mb-4">Add Expense</Button>
      )}
      <div className="w-full max-w-lg mt-4">
        {expenses.length > 0 && (
          <Card className="shadow-lg">
            <CardContent>
              <h2 className="text-xl font-bold mb-2">Expenses</h2>
              <ul className="space-y-2">
                {expenses.map((expense, index) => (
                  <li key={index} className="border-b py-2">
                    <p className="font-medium">{expense.title} - ${expense.price}</p>
                    <p className="text-sm text-gray-600">{expense.description}</p>
                    <p className="text-xs text-gray-500">{expense.date}</p>
                  </li>
                ))}
              </ul>
            </CardContent>
          </Card>
        )}
      </div>
    </div>
  );
};

export default Home;
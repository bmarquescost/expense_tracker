import { useForm } from "react-hook-form";

// UI Components
export function Card({ children, className }) {
  return <div className={`bg-white p-6 rounded-lg ${className}`}>{children}</div>;
}

export function CardContent({ children }) {
  return <div className="p-4">{children}</div>;
}

export function Button({ children, type = "button", className }) {
  return (
    <button type={type} className={`bg-blue-500 text-white p-2 rounded ${className}`}>
      {children}
    </button>
  );
}

export function Input({ id, type = "text", ...props }) {
  return <input id={id} type={type} className="border p-2 w-full rounded" {...props} />;
}

export function Textarea({ id, ...props }) {
  return <textarea id={id} className="border p-2 w-full rounded" {...props} />;
}

export function Label({ htmlFor, children }) {
  return <label htmlFor={htmlFor} className="block font-medium mb-1">{children}</label>;
}

export default function TicketForm() {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm();

  const onSubmit = (data) => {
    console.log("Submitted Data:", data);
  };

  return (
    <Card className="w-full max-w-lg shadow-lg">
      <CardContent>
        <h2 className="text-2xl font-bold mb-4">Expense Form</h2>
        <form onSubmit={handleSubmit(onSubmit)} className="space-y-4">
          <div>
            <Label htmlFor="title">Title</Label>
            <Input id="title" {...register("title", { required: "Title is required" })} />
            {errors.title && <p className="text-red-500 text-sm">{errors.title.message}</p>}
          </div>
          <div>
            <Label htmlFor="price">Price</Label>
            <Input id="price" type="number" step="0.01" {...register("price", { required: "Price is required", valueAsNumber: true })} />
            {errors.price && <p className="text-red-500 text-sm">{errors.price.message}</p>}
          </div>
          <div>
            <Label htmlFor="description">Description</Label>
            <Textarea id="description" {...register("description", { required: "Description is required" })} />
            {errors.description && <p className="text-red-500 text-sm">{errors.description.message}</p>}
          </div>
          <div>
            <Label htmlFor="date">Date</Label>
            <Input id="date" type="date" {...register("date", { required: "Date is required" })} />
            {errors.date && <p className="text-red-500 text-sm">{errors.date.message}</p>}
          </div>
          <Button type="submit" className="w-full">Submit</Button>
        </form>
      </CardContent>
    </Card>
  );
}

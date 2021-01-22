const operationMap = {
  "+": "add",
  "-": "subtract",
  "x": "multiply",
  "÷": "divide",
  "root": "sqrt"
};

export default async function operate(operandOne, operandTwo, operationSymbol) {

  operandOne = operandOne || "0";
  operandTwo = operandTwo || (operationSymbol === "÷" || operationSymbol === 'x' ? "1" : "0"); //If dividing or multiplying, then 1 maintains current value in cases of null

  const operation = operationMap[operationSymbol];
  console.log(`Calling ${operation} service`);

  const rawResponse = await fetch(`/calculate/${operation}`, {
    method: 'POST',
    headers: {
      'Accept': 'application/json',
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({
      operandOne,
      operandTwo
    }),
  });
  console.log("after rawResponse" + rawResponse)
  const response = await rawResponse.json();
  console.log("after response" + response)

  return response.toString();
  
}
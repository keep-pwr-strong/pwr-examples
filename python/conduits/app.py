from flask import Flask, jsonify
from transaction import Transactions
from sync_messages import sync
import threading

app = Flask(__name__)

# Add sync to fetch messages and add it to the pending txs
threading.Thread(target=sync, daemon=True).start()

# Define an HTTP GET route at '/pendingVmTransactions'
# When accessed, this route will return the list of pending transactions
@app.route('/pendingVmTransactions/', methods=['GET'])
def pending_vm_transactions():
    # Retrieve the list of pending transactions using the getPendingTransactions method
    pending_transactions = Transactions.get_pending_transactions()

    array = []
    # Map through each transaction in the pendingTransactions array
    for txn in pending_transactions:
        # Convert each transaction (assumed to be a Buffer or Uint8Array) to a hexadecimal string
        hex_string = '0x' + ''.join(format(byte, '02x') for byte in txn)
        # Return the hexadecimal representation of the transaction
        Transactions.remove(txn)
        array.append(hex_string)
    # Send the resulting array of hex strings as a JSON response
    return jsonify(array), 200

# Set the port number for the server to listen on
# Start the Flask server and listen for connections on the specified port
if __name__ == '__main__':
    port = 8000
    app.run(host='0.0.0.0', port=port, debug=False)
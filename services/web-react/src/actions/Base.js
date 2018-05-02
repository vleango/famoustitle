export let ROOT_API_URL;

if(process.env.NODE_ENV === "production") {
    ROOT_API_URL = `https://u10yz1nzwc.execute-api.us-west-2.amazonaws.com/production`;
} else {
    ROOT_API_URL = `http://localhost:4000`
}

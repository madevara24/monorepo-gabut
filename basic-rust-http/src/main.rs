use axum::{
    routing::get,
    Router,
};
use std::net::SocketAddr;

#[tokio::main]
async fn main() {
    // Build our application with a route
    let app = Router::new().route("/", get(handler));

    // Define the address to listen on
    let addr = SocketAddr::from(([127, 0, 0, 1], 3000));

    println!("Listening on http://{}", addr);

    // Run the server
    axum::Server::bind(&addr)
        .serve(app.into_make_service())
        .await
        .unwrap();
}

// Handler for the root path
async fn handler() -> &'static str {
    "Hello, World!"
}

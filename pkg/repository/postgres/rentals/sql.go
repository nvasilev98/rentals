package rentals

const selectRentalByID = `SELECT 
								r.id, name, description, type, vehicle_make, vehicle_model, vehicle_year,
								vehicle_length, sleeps, primary_image_url, price_per_day, home_city, home_state,
								home_zip, home_country, lat, lng, user_id, first_name, last_name
								FROM rentals r
								LEFT JOIN users u
								ON r.user_id = u.id 
								WHERE r.id = $1`

package datadb

var GetBreedingMales = "SELECT * FROM breeding WHERE male1_suid = ? UNION SELECT * FROM breeding WHERE male2_suid = ? UNION SELECT * FROM breeding WHERE male3_suid = ? UNION SELECT * FROM breeding WHERE male4_suid = ?"

var GetBreedingUuid = "SELECT breeding_uuid FROM breeding WHERE owner_uuid = ? AND suid = ? AND breeding_year = ?"

var GetAllBreeding = "SELECT breeding_year FROM breeding WHERE owner_uuid = ? AND suid = ? ORDER BY breeding_year DESC"

var AddBreedingEvent = "INSERT INTO breeding (uuid, female_suid, male2_suid, male3_suid, male4_suid, female_sid, male1_sid, male2_sid, male3_sid, male4_sid, breeding_year, female_weight, male1_weight, male2_weight, male3_weight, male4_weight, cooling_start, cooling_end, warming_start, warming_end, pairing1_date, pairing2_date, pairing3_date, pairing4_date, gravid_date,lay_date, clutch_size, clutch_survive, outcome, notes) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,?, ?, ?, ?, ?, ?, ?, ?,)"

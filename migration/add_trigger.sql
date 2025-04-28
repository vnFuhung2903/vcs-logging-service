CREATE OR REPLACE FUNCTION log_user_insert()
RETURNS TRIGGER AS $$
BEGIN
    INSERT INTO logs (user_id, operation)
    VALUES (NEW.id, 'INSERT');
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER log_user_insert_trigger
AFTER INSERT ON users
FOR EACH ROW
EXECUTE FUNCTION log_user_insert();


CREATE OR REPLACE FUNCTION log_user_delete()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.deleted_at IS NOT NULL AND OLD.deleted_at IS NULL THEN
        INSERT INTO logs (user_id, operation)
        VALUES (NEW.id, 'DELETE');
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER log_user_delete_trigger
AFTER UPDATE ON users
FOR EACH ROW
EXECUTE FUNCTION log_user_delete();


CREATE OR REPLACE FUNCTION log_user_update()
RETURNS TRIGGER AS $$
BEGIN
    IF OLD.email IS DISTINCT FROM NEW.email THEN
        INSERT INTO logs (user_id, operation, collumn, old_data, new_data)
        VALUES (NEW.id, 'UPDATE', 'email', OLD.email, NEW.email);
    END IF;

    IF OLD.password IS DISTINCT FROM NEW.password THEN
        INSERT INTO logs (user_id, operation, collumn, old_data, new_data)
        VALUES (NEW.id, 'UPDATE', 'password', OLD.password, NEW.password);
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER log_user_update_trigger
AFTER UPDATE ON users
FOR EACH ROW
EXECUTE FUNCTION log_user_update();
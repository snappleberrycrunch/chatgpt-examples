def calculate_DER(putouts, assists, total_chances):
    """
    Calculate Defensive Efficiency Rating (DER) for a first baseman.
    
    DER = (putouts + assists) / total_chances
    
    Args:
    putouts (int): Number of putouts by the first baseman.
    assists (int): Number of assists by the first baseman.
    total_chances (int): Total defensive chances (putouts + assists + errors) for the first baseman.
    
    Returns:
    float: Defensive Efficiency Rating (DER).
    """
    return (putouts + assists) / total_chances
